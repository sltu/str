package str

import "strings"

// ReplaceStart replaces old with repl at the start of the string, if present.
// @group Replace
//
// Example: replace prefix
//
//	v := str.Of("prefix-value").ReplaceStart("prefix-", "new-").String()
//	str.Dump(v)
//	// #string new-value
func (s String) ReplaceStart(old, repl string) String {
	if old == "" || !strings.HasPrefix(s.s, old) {
		return s
	}
	return String{s: repl + s.s[len(old):]}
}
