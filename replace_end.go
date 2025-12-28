package str

import "strings"

// ReplaceEnd replaces old with repl at the end of the string, if present.
// @group Replace
//
// Example: replace suffix
//
//	v := str.Of("file.old").ReplaceEnd(".old", ".new").String()
//	str.Dump(v)
//	// #string file.new
func (s String) ReplaceEnd(old, repl string) String {
	if old == "" || !strings.HasSuffix(s.s, old) {
		return s
	}
	return String{s: s.s[:len(s.s)-len(old)] + repl}
}
