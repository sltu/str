package str

import (
	"strings"
	"unicode"
)

// Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim whitespace
//
//	v := str.Of("  GoForj  ").Trim("").String()
//	str.Dump(v)
//	// #string GoForj
func (s String) Trim(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.Trim(s.s, cutset)}
}

// TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim left
//
//	v := str.Of("  GoForj  ").TrimLeft("").String()
//	str.Dump(v)
//	// #string GoForj
func (s String) TrimLeft(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimLeftFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimLeft(s.s, cutset)}
}

// TrimRight trims trailing characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim right
//
//	v := str.Of("  GoForj  ").TrimRight("").String()
//	str.Dump(v)
//	// #string   GoForj
func (s String) TrimRight(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimRightFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimRight(s.s, cutset)}
}
