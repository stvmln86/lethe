// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a file body string with trimmed whitespace and a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a lowercase file extension string with trimmed whitespace and a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a lowercase alphanumeric file name string with trimmed whitespace.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.In(char, unicode.L, unicode.Nd):
			chars = append(chars, char)
		case char == ' ' || char == '-' || char == '_':
			chars = append(chars, '-')
		}
	}

	return string(chars)
}

// Path returns a cleaned file path string with trimmed whitespace.
func Path(path string) string {
	path = strings.TrimSpace(path)
	return filepath.Clean(path)
}
