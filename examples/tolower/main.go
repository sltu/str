//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ToLower returns a lowercase copy of the string using Unicode rules.

	// Example: lowercase text
	v := str.Of("GoLang").ToLower().String()
	str.Dump(v)
	// #string golang
}
