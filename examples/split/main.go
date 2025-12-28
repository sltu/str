//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Split splits the string by the given separator.

	// Example: split on comma
	v := str.Of("a,b,c").Split(",")
	str.Dump(v)
	// #[]string [a b c]
}
