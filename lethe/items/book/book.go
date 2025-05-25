// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"os"

	"github.com/stvmln86/lethe/lethe/items/note"
	"github.com/stvmln86/lethe/lethe/tools/file"
	"github.com/stvmln86/lethe/lethe/tools/neat"
	"github.com/stvmln86/lethe/lethe/tools/path"
)

// Book is a directory of plaintext note files.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create creates and returns a new empty Note in the Book.
func (b *Book) Create(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)

	if file.Exists(dest) {
		return nil, fmt.Errorf("cannot create note %q - already exists", name)
	}

	if err := file.Write(dest, "", b.Mode); err != nil {
		return nil, fmt.Errorf("cannot create note %q - file error", name)
	}

	return note.New(dest, b.Mode), nil
}

// Filter returns all Notes in the Book that pass a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note
	for _, note := range b.List() {
		ok, err := ffun(note)
		switch {
		case ok:
			notes = append(notes, note)
		case err != nil:
			return nil, err
		}
	}

	return notes, nil
}

// Get returns an existing Note from the Book.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)

	if !file.Exists(dest) {
		return nil, fmt.Errorf("cannot get note %q - does not exist", name)
	}

	return note.New(dest, b.Mode), nil
}

// GetOrCreate returns a new or existing Note from the Book.
func (b *Book) GetOrCreate(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)

	if !file.Exists(dest) {
		return b.Create(name)
	}

	return b.Get(name)
}

// List returns all Notes in the Book in sorted order.
func (b *Book) List() []*note.Note {
	var notes []*note.Note
	for _, orig := range file.List(b.Dire, b.Extn) {
		note := note.New(orig, b.Mode)
		notes = append(notes, note)
	}

	return notes
}

// Match returns all Notes in the Book with names containing a case-insensitive substring.
func (b *Book) Match(subs string) []*note.Note {
	notes, _ := b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(subs), nil
	})

	return notes
}

// Search returns all Notes in the Book with bodies containing a case-insensitive substring.
func (b *Book) Search(subs string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(subs)
	})
}
