//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Remove deletes all occurrences of provided substrings.

	// Example: remove substrings
	v := str.Of("The Go Toolkit").Remove("Go ").String()
	str.Dump(v)
	// #string The Toolkit
}
