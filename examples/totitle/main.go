//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.

	// Example: title map runes
	v := str.Of("ß").ToTitle().String()
	str.Dump(v)
	// #string SS
}
