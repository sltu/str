package str

import "strings"

// Repeat repeats the string count times (non-negative).
// @group Transform
//
// Example: repeat string
//
//	v := str.Of("go").Repeat(3).String()
//	str.Dump(v)
//	// #string gogogo
func (s String) Repeat(count int) String {
	if count <= 0 {
		return String{s: ""}
	}
	return String{s: strings.Repeat(s.s, count)}
}
