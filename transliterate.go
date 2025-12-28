package str

import "strings"

// Transliterate replaces a small set of accented runes with ASCII equivalents.
// @group Transform
//
// Example: strip accents
//
//	v := str.Of("café déjà vu").Transliterate().String()
//	str.Dump(v)
//	// #string cafe deja vu
func (s String) Transliterate() String {
	if s.s == "" {
		return s
	}

	var b strings.Builder
	b.Grow(len(s.s))
	changed := false

	for _, r := range s.s {
		if mapped, ok := transliterateBasic(r); ok {
			b.WriteString(mapped)
			changed = true
			continue
		}
		b.WriteRune(r)
	}

	if !changed {
		return s
	}
	return String{s: b.String()}
}
