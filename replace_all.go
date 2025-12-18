package str

import "strings"

// ReplaceAll replaces all occurrences of old with new in the string.
// If old is empty, the original string is returned unchanged.
// @group Replace
//
// Example: replace all occurrences
//
//	v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
//	godump.Dump(v)
//	// #string Go Gopher Go
func (s String) ReplaceAll(old, new string) String {
	if old == "" {
		return s
	}
	return String{s: strings.ReplaceAll(s.s, old, new)}
}
