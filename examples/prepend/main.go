//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Prepend concatenates the provided parts to the beginning of the string.

	// Example: prepend text
	v := str.Of("World").Prepend("Hello ", "Go ").String()
	str.Dump(v)
	// #string Hello Go World
}
