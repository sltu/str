package str

import (
	"regexp"
	"sort"
	"strings"
)

// ReplaceFirst replaces the first occurrence of old with repl.
// @group Replace
//
// Example: replace first
//
//	v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
//	str.Dump(v)
//	// #string go gopher
func (s String) ReplaceFirst(old, repl string) String {
	return String{s: strings.Replace(s.s, old, repl, 1)}
}

// ReplaceLast replaces the last occurrence of old with repl.
// @group Replace
//
// Example: replace last
//
//	v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
//	str.Dump(v)
//	// #string gopher go
func (s String) ReplaceLast(old, repl string) String {
	idx := strings.LastIndex(s.s, old)
	if idx == -1 || old == "" {
		return s
	}
	var b strings.Builder
	b.Grow(len(s.s) - len(old) + len(repl))
	b.WriteString(s.s[:idx])
	b.WriteString(repl)
	b.WriteString(s.s[idx+len(old):])
	return String{s: b.String()}
}

// ReplaceArray replaces all occurrences of each old in olds with repl.
// @group Replace
//
// Example: replace many
//
//	v := str.Of("The---Go---Toolkit")
//	str.Dump(v.ReplaceArray([]string{"---"}, "-").String())
//	// #string The-Go-Toolkit
func (s String) ReplaceArray(olds []string, repl string) String {
	out := s.s
	for _, old := range olds {
		if old == "" {
			continue
		}
		out = strings.ReplaceAll(out, old, repl)
	}
	return String{s: out}
}

// ReplaceMatches applies repl to each regex match and returns the result.
// @group Replace
//
// Example: regex replace with callback
//
//	re := regexp.MustCompile(`\d+`)
//	v := str.Of("Hello 123 World").ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String()
//	str.Dump(v)
//	// #string Hello [123] World
func (s String) ReplaceMatches(pattern *regexp.Regexp, repl func(string) string) String {
	return String{s: pattern.ReplaceAllStringFunc(s.s, repl)}
}

// Swap replaces multiple values using strings.Replacer built from a map.
// @group Replace
//
// Example: swap map
//
//	pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
//	v := str.Of("Gophers are great!").Swap(pairs).String()
//	str.Dump(v)
//	// #string GoForj is fantastic!
func (s String) Swap(pairs map[string]string) String {
	if len(pairs) == 0 {
		return s
	}
	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

	repPairs := make([]string, 0, len(keys)*2)
	for _, k := range keys {
		repPairs = append(repPairs, k, pairs[k])
	}

	r := strings.NewReplacer(repPairs...)
	return String{s: r.Replace(s.s)}
}
