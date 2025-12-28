package str

import "strings"

// Join joins the provided words with sep.
// @group Words
//
// Example: join words
//
//	v := str.Of("unused").Join([]string{"foo", "bar"}, "-").String()
//	str.Dump(v)
//	// #string foo-bar
func (s String) Join(words []string, sep string) String {
	return String{s: strings.Join(words, sep)}
}
