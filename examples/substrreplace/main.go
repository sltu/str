//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// SubstrReplace replaces the rune slice in [start:end) with repl.

	// Example: replace range
	v := str.Of("naïve café").SubstrReplace("i", 2, 3).String()
	str.Dump(v)
	// #string naive café
}
