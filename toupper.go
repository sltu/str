package str

import "strings"

// ToUpper returns an uppercase copy of the string using Unicode rules.
// @group Case
//
// Example: uppercase text
//
//	v := str.Of("GoLang").ToUpper().String()
//	str.Dump(v)
//	// #string GOLANG
func (s String) ToUpper() String {
	return String{s: strings.ToUpper(s.s)}
}
