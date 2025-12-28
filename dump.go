package str

import "github.com/goforj/godump"

// Dump is a thin alias for godump.Dump, provided to keep examples concise
// and readable in documentation and code snippets.
func Dump(vs ...any) {
	godump.Dump(vs...)
}
