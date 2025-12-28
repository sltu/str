package str

import (
	"strings"
	"unicode/utf8"
)

// ReplaceFold replaces all occurrences of old with repl using Unicode case folding.
// @group Replace
//
// Example: replace all (case-insensitive)
//
//	v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
//	str.Dump(v)
//	// #string Go Gopher Go
func (s String) ReplaceFold(old, repl string) String {
	if old == "" {
		return s
	}
	out, ok := replaceFoldAll(s.s, old, repl)
	if !ok {
		return s
	}
	return String{s: out}
}

// ReplaceFirstFold replaces the first occurrence of old with repl using Unicode case folding.
// @group Replace
//
// Example: replace first (case-insensitive)
//
//	v := str.Of("go gopher GO").ReplaceFirstFold("GO", "Go").String()
//	str.Dump(v)
//	// #string Go gopher GO
func (s String) ReplaceFirstFold(old, repl string) String {
	if old == "" {
		return s
	}
	start, end, ok := findFoldMatch(s.s, old, false)
	if !ok {
		return s
	}
	return String{s: replaceAt(s.s, start, end, repl)}
}

// ReplaceLastFold replaces the last occurrence of old with repl using Unicode case folding.
// @group Replace
//
// Example: replace last (case-insensitive)
//
//	v := str.Of("go gopher GO").ReplaceLastFold("GO", "Go").String()
//	str.Dump(v)
//	// #string go gopher Go
func (s String) ReplaceLastFold(old, repl string) String {
	if old == "" {
		return s
	}
	start, end, ok := findFoldMatch(s.s, old, true)
	if !ok {
		return s
	}
	return String{s: replaceAt(s.s, start, end, repl)}
}

func replaceFoldAll(s, old, repl string) (string, bool) {
	runeStarts, oldLen := foldRuneStarts(s, old)
	if oldLen == 0 || len(runeStarts)-1 < oldLen {
		return s, false
	}

	var b strings.Builder
	b.Grow(len(s))
	matched := false
	lastByte := 0

	for i := 0; i+oldLen <= len(runeStarts)-1; {
		start := runeStarts[i]
		end := runeStarts[i+oldLen]
		if strings.EqualFold(s[start:end], old) {
			b.WriteString(s[lastByte:start])
			b.WriteString(repl)
			lastByte = end
			matched = true
			i += oldLen
			continue
		}
		i++
	}

	if !matched {
		return s, false
	}

	b.WriteString(s[lastByte:])
	return b.String(), true
}

func findFoldMatch(s, old string, last bool) (int, int, bool) {
	runeStarts, oldLen := foldRuneStarts(s, old)
	if oldLen == 0 || len(runeStarts)-1 < oldLen {
		return 0, 0, false
	}

	found := false
	matchStart := 0
	matchEnd := 0

	for i := 0; i+oldLen <= len(runeStarts)-1; i++ {
		start := runeStarts[i]
		end := runeStarts[i+oldLen]
		if strings.EqualFold(s[start:end], old) {
			matchStart = start
			matchEnd = end
			found = true
			if !last {
				return matchStart, matchEnd, true
			}
		}
	}

	if !found {
		return 0, 0, false
	}
	return matchStart, matchEnd, true
}

func foldRuneStarts(s, old string) ([]int, int) {
	runeStarts := make([]int, 0, utf8.RuneCountInString(s)+1)
	for i := range s {
		runeStarts = append(runeStarts, i)
	}
	runeStarts = append(runeStarts, len(s))
	return runeStarts, len([]rune(old))
}

func replaceAt(s string, start, end int, repl string) string {
	var b strings.Builder
	b.Grow(len(s) - (end - start) + len(repl))
	b.WriteString(s[:start])
	b.WriteString(repl)
	b.WriteString(s[end:])
	return b.String()
}
