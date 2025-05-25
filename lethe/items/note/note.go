// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"os"

	"github.com/stvmln86/lethe/lethe/tools/file"
	"github.com/stvmln86/lethe/lethe/tools/neat"
	"github.com/stvmln86/lethe/lethe/tools/path"
)

// Note is a single plaintext note file in a directory.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete deletes the Note.
func (n *Note) Delete() error {
	if !n.Exists() {
		return fmt.Errorf("cannot delete note %q - does not exist", n.Name())
	}

	if err := file.Delete(n.Orig); err != nil {
		return fmt.Errorf("cannot delete note %q - file error", n.Name())
	}

	return nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name contains a case-insensitive substring.
func (n *Note) Match(subs string) bool {
	return neat.Substring(n.Name(), subs)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	if !n.Exists() {
		return "", fmt.Errorf("cannot read note %q - does not exist", n.Name())
	}

	body, err := file.Read(n.Orig)
	if err != nil {
		return "", fmt.Errorf("cannot read note %q - file error", n.Name())
	}

	return neat.Body(body), nil
}

// Rename renames the Note to a different name.
func (n *Note) Rename(name string) error {
	if !n.Exists() {
		return fmt.Errorf("cannot rename note %q - does not exist", n.Name())
	}

	name = neat.Name(name)
	if err := file.Rename(n.Orig, "name", name); err != nil {
		return fmt.Errorf("cannot rename note %q - file error", n.Name())
	}

	return nil
}

// Search returns true if the Note's body contains a case-insensitive substring.
func (n *Note) Search(subs string) (bool, error) {
	if !n.Exists() {
		return false, fmt.Errorf("cannot search note %q - does not exist", n.Name())
	}

	body, err := file.Read(n.Orig)
	if err != nil {
		return false, fmt.Errorf("cannot search note %q - file error", n.Name())
	}

	return neat.Substring(body, subs), nil
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	if !n.Exists() {
		return fmt.Errorf("cannot update note %q - does not exist", n.Name())
	}

	body = neat.Body(body)
	if err := file.Write(n.Orig, body, n.Mode); err != nil {
		return fmt.Errorf("cannot update note %q - file error", n.Name())
	}

	return nil
}
