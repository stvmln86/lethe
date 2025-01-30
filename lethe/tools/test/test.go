// Package test implements unit testing data and functions.
package test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a base:body map of mock note files.
var MockFiles = map[string]string{
	"alpha.extn":   "Alpha note.\n",
	"bravo.extn":   "Bravo note.\n",
	"charlie.extn": "Charlie note.\n",
	"delta.trash":  "Delta note (trash).\n",
}

// AssertErr asserts an error format string matches an error message.
func AssertErr(t *testing.T, err error, text string) {
	text = regexp.MustCompile(`%q`).ReplaceAllString(text, `".*"`)
	text = regexp.MustCompile(`%\w`).ReplaceAllString(text, `.*`)
	assert.Regexp(t, text, err.Error())
}

// AssertDire asserts a directory's files are equal to a base:body map.
func AssertDire(t *testing.T, dire string, files map[string]string) {
	for base, body := range files {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, body)
	}
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// TempDire returns a temporary directory populated with MockFiles entries.
func TempDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockFiles {
		dest := filepath.Join(dire, base)
		os.WriteFile(dest, []byte(body), 0666)
	}

	return dire
}

// TempFile returns a temporary file from a MockFiles entry.
func TempFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	os.WriteFile(dest, []byte(MockFiles[base]), 0666)
	return dest
}
