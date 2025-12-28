//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Squish trims leading/trailing whitespace and collapses internal whitespace to single spaces.

	// Example: squish whitespace
	v := str.Of("   go   forj  ").Squish().String()
	str.Dump(v)
	// #string go forj
}
