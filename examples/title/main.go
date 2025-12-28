//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.

	// Example: title case words
	v := str.Of("a nice title uses the correct case").Title().String()
	str.Dump(v)
	// #string A Nice Title Uses The Correct Case
}
