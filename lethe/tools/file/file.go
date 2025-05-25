// Package file implements file system handling functions.
package file

import (
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/stvmln86/lethe/lethe/tools/path"
)

// Delete deletes an existing file.
func Delete(orig string) error {
	return os.Remove(orig)
}

// Exists returns true if a file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// List returns a sorted slice of all files in a directory with an extension.
func List(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	slices.Sort(origs)
	return origs
}

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	bytes, err := os.ReadFile(orig)
	return string(bytes), err
}

// Rename renames an existing file to a different directory, name or extension.
func Rename(orig, mode, swap string) error {
	dest := path.Rename(orig, mode, swap)
	return os.Rename(orig, dest)
}

// Write overwrites a new or existing file with a string.
func Write(orig, body string, mode os.FileMode) error {
	return os.WriteFile(orig, []byte(body), mode)
}
