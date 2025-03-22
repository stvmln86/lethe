// Package path implements file path creation and manipulation functions.
package path

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Extn returns a path's file extension string without a leading dot.
func Extn(orig string) string {
	base := filepath.Base(orig)
	return append(strings.SplitN(base, ".", 2), "")[1]
}

// Join returns a path from a directory, slug and extension string.
func Join(dire, slug, extn string) string {
	return filepath.Join(dire, fmt.Sprintf("%s.%s", slug, extn))
}

// Match returns true if a path's slug contains a case-insensitive substring.
func Match(orig, subs string) bool {
	subs = strings.ToLower(subs)
	slug := strings.ToLower(Slug(orig))
	return strings.Contains(slug, subs)
}

// Reextn returns a path with a different extension string.
func Reextn(orig, extn string) string {
	dire := filepath.Dir(orig)
	return Join(dire, Slug(orig), extn)
}

// Reslug returns a path with a different slug string.
func Reslug(orig, slug string) string {
	dire := filepath.Dir(orig)
	return Join(dire, slug, Extn(orig))
}

// Slug returns a path's file slug string.
func Slug(orig string) string {
	base := filepath.Base(orig)
	return strings.SplitN(base, ".", 2)[0]
}
