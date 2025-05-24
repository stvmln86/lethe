// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a trimmed file body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a trimmed lowercase file extension string with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a trimmed lowercase alphanumeric file name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char), unicode.IsNumber(char):
			chars = append(chars, char)
		case unicode.IsSpace(char), char == '-', char == '_':
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}

// Path returns a trimmed normalised file path string.
func Path(path string) string {
	path = strings.TrimSpace(path)
	return filepath.Clean(path)
}
