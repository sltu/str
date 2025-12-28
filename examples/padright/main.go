//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// PadRight pads the string on the right to reach length runes using pad (defaults to space).

	// Example: pad right
	v := str.Of("go").PadRight(5, ".").String()
	str.Dump(v)
	// #string go...
}
