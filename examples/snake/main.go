//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Snake converts the string to snake_case using the provided separator (default "_").

	// Example: snake case
	v := str.Of("fooBar baz").Snake("_").String()
	str.Dump(v)
	// #string foo_bar_baz
}
