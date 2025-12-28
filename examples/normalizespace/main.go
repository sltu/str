//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// NormalizeSpace collapses whitespace runs to single spaces without trimming ends.

	// Example: normalize interior space
	v := str.Of("  go   forj  ").NormalizeSpace().String()
	str.Dump(v)
	// #string  go forj 
}
