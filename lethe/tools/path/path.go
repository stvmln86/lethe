// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Base returns a path's base name with the file extension.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Extn returns a path's file extension with the leading dot.
func Extn(orig string) string {
	_, extn, _ := strings.Cut(orig, ".")
	return "." + extn
}

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Glob returns all paths in a directory with a file extension.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	return origs
}

// Join returns a joined path from a directory, name and file extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Name returns a path's base name without the file extension.
func Name(orig string) string {
	base := filepath.Base(orig)
	return strings.SplitN(base, ".", 2)[0]
}

// Reextn returns a path with a changed file extension.
func Reextn(orig, extn string) string {
	dire := filepath.Dir(orig)
	name := Name(orig)
	return filepath.Join(dire, name+extn)
}

// Rename returns a path with a changed file extension.
func Rename(orig, name string) string {
	dire := filepath.Dir(orig)
	extn := Extn(orig)
	return filepath.Join(dire, name+extn)
}
