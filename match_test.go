package str

import (
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`g(o+)`)
	val := Of("gooo")
	got := val.Match(re)
	if len(got) != 2 || got[0] != "gooo" || got[1] != "ooo" {
		t.Fatalf("Match = %#v", got)
	}

	if got := val.Match(nil); got != nil {
		t.Fatalf("Match nil = %#v", got)
	}
}

func TestMatchAll(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`go+`)
	got := Of("go gopher gooo").MatchAll(re)
	if len(got) != 3 || got[0][0] != "go" || got[1][0] != "go" || got[2][0] != "gooo" {
		t.Fatalf("MatchAll = %#v", got)
	}

	if got := Of("none").MatchAll(re); got != nil {
		t.Fatalf("MatchAll none = %#v", got)
	}
	if got := Of("none").MatchAll(nil); got != nil {
		t.Fatalf("MatchAll nil = %#v", got)
	}
}
