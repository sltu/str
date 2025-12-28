//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/str"
	"regexp"
)

func main() {
	// ReplaceMatches applies repl to each regex match and returns the result.

	// Example: regex replace with callback
	re := regexp.MustCompile(`\d+`)
	v := str.Of("Hello 123 World").ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String()
	str.Dump(v)
	// #string Hello [123] World
}
