package str

import "testing"

func TestReplaceStartEnd(t *testing.T) {
	t.Parallel()

	val := Of("prefix-value")
	if got := val.ReplaceStart("prefix-", "new-").String(); got != "new-value" {
		t.Fatalf("ReplaceStart = %q", got)
	}
	if got := val.ReplaceStart("missing", "new-").String(); got != "prefix-value" {
		t.Fatalf("ReplaceStart missing = %q", got)
	}
	if got := val.ReplaceStart("", "new-").String(); got != "prefix-value" {
		t.Fatalf("ReplaceStart empty old = %q", got)
	}

	val = Of("file.old")
	if got := val.ReplaceEnd(".old", ".new").String(); got != "file.new" {
		t.Fatalf("ReplaceEnd = %q", got)
	}
	if got := val.ReplaceEnd(".missing", ".new").String(); got != "file.old" {
		t.Fatalf("ReplaceEnd missing = %q", got)
	}
	if got := val.ReplaceEnd("", ".new").String(); got != "file.old" {
		t.Fatalf("ReplaceEnd empty old = %q", got)
	}
}
