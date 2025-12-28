//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// UcFirst returns the string with the first rune upper-cased.

	// Example: uppercase first rune
	v := str.Of("gopher").UcFirst().String()
	str.Dump(v)
	// #string Gopher
}
