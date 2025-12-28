//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Camel converts the string to camelCase.

	// Example: camel case
	v := str.Of("foo_bar baz").Camel().String()
	str.Dump(v)
	// #string fooBarBaz
}
