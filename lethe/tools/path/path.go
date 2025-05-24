// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Base returns a path's base name.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's file extension with a leading dot.
func Extn(orig string) string {
	base := Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[clip:]
	}

	return "."
}

// Join returns a path joined from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Name returns a path's base name without the extension.
func Name(orig string) string {
	base := Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return base
}

// Rename returns a path with a changed directory, name or extension.
func Rename(orig, mode, swap string) string {
	switch mode {
	case "dire":
		return Join(swap, Name(orig), Extn(orig))
	case "name":
		return Join(Dire(orig), swap, Extn(orig))
	case "extn":
		return Join(Dire(orig), Name(orig), swap)
	default:
		panic("invalid rename mode")
	}
}
