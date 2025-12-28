//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// PadBoth pads the string on both sides to reach length runes using pad (defaults to space).

	// Example: pad both
	v := str.Of("go").PadBoth(6, "-").String()
	str.Dump(v)
	// #string --go--
}
