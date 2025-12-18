//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ReplaceAll replaces all occurrences of old with new in the string.
	// If old is empty, the original string is returned unchanged.

	// Example: replace all occurrences
	v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
	godump.Dump(v)
	// #string Go Gopher Go
}
