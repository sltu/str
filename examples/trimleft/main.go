//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.

	// Example: trim left
	v := str.Of("  GoForj  ").TrimLeft("").String()
	str.Dump(v)
	// #string GoForj
}
