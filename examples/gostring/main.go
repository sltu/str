//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/str"
)

func main() {
	// GoString allows %#v formatting to print the raw string.

	// Example: fmt %#v uses GoString
	v := str.Of("go")
	str.Dump(fmt.Sprintf("%#v", v))
	// #string go
}
