package book

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/items/note"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func mockBook(t *testing.T) *Book {
	dire := test.MockDire(t)
	return New(dire, ".extn", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := mockBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Create("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.extn"), note.Orig)
	assert.FileExists(t, note.Orig)
	assert.NoError(t, err)

	// error - already exists
	note, err = book.Create("test")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot create note "test" - already exists`)
}

func TestFilter(t *testing.T) {
	// setup
	book := mockBook(t)
	ffun := func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	}

	// success
	notes, err := book.Filter(ffun)
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, filepath.Join(book.Dire, "alpha.extn"), note.Orig)
	assert.NoError(t, err)

	// error - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot get note "nope" - does not exist`)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success - create
	note, err := book.GetOrCreate("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.extn"), note.Orig)
	assert.NoError(t, err)

	// success - get
	note, err = book.GetOrCreate("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.extn"), note.Orig)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.Equal(t, "bravo", notes[1].Name())
}

func TestMatch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.Match("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
}

func TestSearch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Search("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}
