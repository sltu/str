//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// SplitWords splits the string into words (Unicode letters/digits runs).

	// Example: split words
	v := str.Of("one, two, three").SplitWords()
	str.Dump(v)
	// #[]string [one two three]
}
