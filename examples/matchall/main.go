//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/str"
	"regexp"
)

func main() {
	// MatchAll returns all matches and submatches for the pattern.

	// Example: regex match all
	re := regexp.MustCompile(`go+`)
	v := str.Of("go gopher gooo").MatchAll(re)
	str.Dump(v)
	// #[][]string [[go] [gooo]]
}
