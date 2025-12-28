package str

import "strings"

// ToLower returns a lowercase copy of the string using Unicode rules.
// @group Case
//
// Example: lowercase text
//
//	v := str.Of("GoLang").ToLower().String()
//	str.Dump(v)
//	// #string golang
func (s String) ToLower() String {
	return String{s: strings.ToLower(s.s)}
}
