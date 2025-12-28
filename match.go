package str

import "regexp"

// Match returns the first match and submatches for the pattern.
// @group Match
//
// Example: regex match
//
//	re := regexp.MustCompile(`g(o+)`)
//	v := str.Of("gooo").Match(re)
//	str.Dump(v)
//	// #[]string [gooo ooo]
func (s String) Match(pattern *regexp.Regexp) []string {
	if pattern == nil {
		return nil
	}
	return pattern.FindStringSubmatch(s.s)
}
