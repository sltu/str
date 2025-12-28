//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceFold replaces all occurrences of old with repl using Unicode case folding.

	// Example: replace all (case-insensitive)
	v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
	str.Dump(v)
	// #string Go Gopher Go
}
