//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// PadLeft pads the string on the left to reach length runes using pad (defaults to space).

	// Example: pad left
	v := str.Of("go").PadLeft(5, " ").String()
	str.Dump(v)
	// #string \u00a0\u00a0\u00a0go
}
