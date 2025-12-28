package str

// Index returns the rune index of the first occurrence of sub, or -1 if not found.
// @group Search
//
// Example: first rune index
//
//	v := str.Of("héllo").Index("llo")
//	str.Dump(v)
//	// #int 2
func (s String) Index(sub string) int {
	return runeIndex(s.s, sub, false)
}

// LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.
// @group Search
//
// Example: last rune index
//
//	v := str.Of("go gophers go").LastIndex("go")
//	str.Dump(v)
//	// #int 10
func (s String) LastIndex(sub string) int {
	return runeIndex(s.s, sub, true)
}

// Slice returns the substring between rune offsets [start:end).
// Indices are clamped; if start >= end the result is empty.
// @group Substrings
//
// Example: rune-safe slice
//
//	v := str.Of("naïve café").Slice(3, 7).String()
//	str.Dump(v)
//	// #string e ca
func (s String) Slice(start, end int) String {
	runes := []rune(s.s)
	start, end = clampRange(start, end, len(runes))
	if start >= end {
		return String{s: ""}
	}
	return String{s: string(runes[start:end])}
}
