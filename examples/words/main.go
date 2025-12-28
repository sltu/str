//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Words limits the string to count words, appending suffix if truncated.

	// Example: limit words
	v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
	str.Dump(v)
	// #string Perfectly balanced as >>>
}
