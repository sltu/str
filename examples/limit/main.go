//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Limit truncates the string to length runes, appending suffix if truncation occurs.

	// Example: limit with suffix
	v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
	str.Dump(v)
	// #string Perfectly...
}
