package str

import "regexp"

// MatchAll returns all matches and submatches for the pattern.
// @group Match
//
// Example: regex match all
//
//	re := regexp.MustCompile(`go+`)
//	v := str.Of("go gopher gooo").MatchAll(re)
//	str.Dump(v)
//	// #[][]string [[go] [gooo]]
func (s String) MatchAll(pattern *regexp.Regexp) [][]string {
	if pattern == nil {
		return nil
	}
	return pattern.FindAllStringSubmatch(s.s, -1)
}
