//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Equals reports whether the string exactly matches other (case-sensitive).

	// Example: exact match
	v := str.Of("gopher").Equals("gopher")
	str.Dump(v)
	// #bool true
}
