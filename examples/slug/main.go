//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
	// Not locale-aware; intended for identifiers/URLs.

	// Example: build slug
	v := str.Of("Go Forj Toolkit").Slug("-").String()
	str.Dump(v)
	// #string go-forj-toolkit
}
