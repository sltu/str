package str

// LastWord returns the last word token or empty string.
// @group Words
//
// Example: last word
//
//	v := str.Of("Hello world").LastWord().String()
//	str.Dump(v)
//	// #string world
func (s String) LastWord() String {
	words := splitWordsRunes(s.s)
	if len(words) == 0 {
		return String{s: ""}
	}
	return String{s: words[len(words)-1]}
}
