//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.

	// Example: trim whitespace
	v := str.Of("  GoForj  ").Trim("").String()
	str.Dump(v)
	// #string GoForj
}
