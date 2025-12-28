package str

// SubstrReplace replaces the rune slice in [start:end) with repl.
// @group Substrings
//
// Example: replace range
//
//	v := str.Of("naïve café").SubstrReplace("i", 2, 3).String()
//	str.Dump(v)
//	// #string naive café
func (s String) SubstrReplace(repl string, start, end int) String {
	runes := []rune(s.s)
	start, end = clampRange(start, end, len(runes))
	if start > end {
		return s
	}
	out := append([]rune{}, runes[:start]...)
	out = append(out, []rune(repl)...)
	out = append(out, runes[end:]...)
	return String{s: string(out)}
}
