//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Pascal converts the string to PascalCase.

	// Example: pascal case
	v := str.Of("foo_bar baz")
	str.Dump(v)
	// #string FooBarBaz
}
