//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/str"
	"regexp"
)

func main() {
	// Match returns the first match and submatches for the pattern.

	// Example: regex match
	re := regexp.MustCompile(`g(o+)`)
	v := str.Of("gooo").Match(re)
	str.Dump(v)
	// #[]string [gooo ooo]
}
