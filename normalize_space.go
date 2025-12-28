package str

import "unicode"

// NormalizeSpace collapses whitespace runs to single spaces without trimming ends.
// @group Cleanup
//
// Example: normalize interior space
//
//	v := str.Of("  go   forj  ").NormalizeSpace().String()
//	str.Dump(v)
//	// #string  go forj 
func (s String) NormalizeSpace() String {
	if s.s == "" {
		return s
	}

	var out []rune
	inSpace := false

	for _, r := range s.s {
		if unicode.IsSpace(r) {
			if !inSpace {
				out = append(out, ' ')
				inSpace = true
			}
			continue
		}
		out = append(out, r)
		inSpace = false
	}

	return String{s: string(out)}
}
