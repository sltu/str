//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceEnd replaces old with repl at the end of the string, if present.

	// Example: replace suffix
	v := str.Of("file.old").ReplaceEnd(".old", ".new").String()
	str.Dump(v)
	// #string file.new
}
