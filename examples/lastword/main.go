//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// LastWord returns the last word token or empty string.

	// Example: last word
	v := str.Of("Hello world").LastWord().String()
	str.Dump(v)
	// #string world
}
