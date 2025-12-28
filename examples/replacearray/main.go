//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceArray replaces all occurrences of each old in olds with repl.

	// Example: replace many
	v := str.Of("The---Go---Toolkit")
	str.Dump(v.ReplaceArray([]string{"---"}, "-").String())
	// #string The-Go-Toolkit
}
