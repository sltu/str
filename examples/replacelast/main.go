//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceLast replaces the last occurrence of old with repl.

	// Example: replace last
	v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
	str.Dump(v)
	// #string gopher go
}
