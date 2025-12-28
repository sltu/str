//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Lines splits the string into lines after normalizing newline variants.

	// Example: split lines
	v := str.Of("a\\r\\nb\\nc").Lines()
	str.Dump(v)
	// #[]string [a b c]
}
