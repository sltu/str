package str

import "strings"

// Remove deletes all occurrences of provided substrings.
// @group Replace
//
// Example: remove substrings
//
//	v := str.Of("The Go Toolkit").Remove("Go ").String()
//	str.Dump(v)
//	// #string The Toolkit
func (s String) Remove(subs ...string) String {
	out := s.s
	for _, sub := range subs {
		if sub == "" {
			continue
		}
		out = strings.ReplaceAll(out, sub, "")
	}
	return String{s: out}
}
