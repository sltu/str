//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// UcWords uppercases the first rune of each word, leaving the rest unchanged.
	// Words are sequences of letters/digits.

	// Example: uppercase each word start
	v := str.Of("hello WORLD").UcWords().String()
	str.Dump(v)
	// #string Hello WORLD
}
