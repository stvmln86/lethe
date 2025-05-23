// Package test implements unit testing data and fixtures.
package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a base:body map of mock file strings.
var MockFiles = map[string]string{
	"alpha.extn":    "Alpha note.\n",
	"bravo.extn":    "Bravo note.\n",
	"charlie.trash": "Charlie note.\n",
	"config.json": `{
		"ext_pattern": "*.*",
		"hard_delete": false,
		"lower_names": true,
		"trim_trails": true,
	}`,
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

// MockDire returns a temporary directory populated with all MockFiles entries.
func MockDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockFiles {
		dest := filepath.Join(dire, base)
		os.WriteFile(dest, []byte(body), 0666)
	}

	return dire
}

// MockFile returns a temporary file populated with a MockFiles entry.
func MockFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	os.WriteFile(dest, []byte(MockFiles[base]), 0666)
	return dest
}
