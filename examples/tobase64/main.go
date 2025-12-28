//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ToBase64 encodes the string using standard Base64.

	// Example: base64 encode
	v := str.Of("gopher").ToBase64().String()
	str.Dump(v)
	// #string Z29waGVy
}
