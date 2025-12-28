//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Transliterate replaces a small set of accented runes with ASCII equivalents.

	// Example: strip accents
	v := str.Of("café déjà vu").Transliterate().String()
	str.Dump(v)
	// #string cafe deja vu
}
