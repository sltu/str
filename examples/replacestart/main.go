//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceStart replaces old with repl at the start of the string, if present.

	// Example: replace prefix
	v := str.Of("prefix-value").ReplaceStart("prefix-", "new-").String()
	str.Dump(v)
	// #string new-value
}
