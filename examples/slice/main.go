//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Slice returns the substring between rune offsets [start:end).
	// Indices are clamped; if start >= end the result is empty.

	// Example: rune-safe slice
	v := str.Of("naïve café").Slice(3, 7).String()
	str.Dump(v)
	// #string e ca
}
