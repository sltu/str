//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Reverse returns a rune-safe reversed string.

	// Example: reverse
	v := str.Of("naïve").Reverse().String()
	str.Dump(v)
	// #string evïan
}
