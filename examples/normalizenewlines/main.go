//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.

	// Example: normalize newline variants
	v := str.Of("a\\r\\nb\\u2028c").NormalizeNewlines().String()
	str.Dump(v)
	// #string a\nb\nc
}
