//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Swap replaces multiple values using strings.Replacer built from a map.

	// Example: swap map
	pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
	v := str.Of("Gophers are great!").Swap(pairs).String()
	str.Dump(v)
	// #string GoForj is fantastic!
}
