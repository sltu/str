//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceFirstFold replaces the first occurrence of old with repl using Unicode case folding.

	// Example: replace first (case-insensitive)
	v := str.Of("go gopher GO").ReplaceFirstFold("GO", "Go").String()
	str.Dump(v)
	// #string Go gopher GO
}
