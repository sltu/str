//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Kebab converts the string to kebab-case.

	// Example: kebab case
	v := str.Of("fooBar baz").Kebab().String()
	str.Dump(v)
	// #string foo-bar-baz
}
