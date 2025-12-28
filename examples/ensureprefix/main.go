//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// EnsurePrefix ensures the string starts with prefix, adding it if missing.

	// Example: ensure prefix
	v := str.Of("path/to").EnsurePrefix("/").String()
	str.Dump(v)
	// #string /path/to
}
