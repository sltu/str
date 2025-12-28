//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Deduplicate collapses consecutive instances of char into a single instance.
	// If char is zero, space is used.

	// Example: collapse spaces
	v := str.Of("The   Go   Playground").Deduplicate(' ').String()
	str.Dump(v)
	// #string The Go Playground
}
