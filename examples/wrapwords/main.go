//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.

	// Example: wrap words
	v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
	str.Dump(v)
	// #string The quick brown fox\njumped over the lazy\ndog.
}
