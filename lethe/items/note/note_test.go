package note

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func mockNote(t *testing.T) *Note {
	orig := test.MockFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := mockNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.NoError(t, err)

	// error - does not exist
	err = New("/nope", 0666).Delete()
	assert.EqualError(t, err, `cannot delete note "nope" - does not exist`)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestMatch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Match("ALPH")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)

	// error - does not exist
	body, err = New("/nope", 0666).Read()
	assert.Empty(t, body)
	assert.EqualError(t, err, `cannot read note "nope" - does not exist`)
}

func TestRename(t *testing.T) {
	// setup
	note := mockNote(t)
	dest := strings.Replace(note.Orig, "alpha", "test", 1)

	// success
	err := note.Rename("test")
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// error - does not exist
	err = New("/nope", 0666).Rename("test")
	assert.EqualError(t, err, `cannot rename note "nope" - does not exist`)
}

func TestSearch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok, err := note.Search("ALPH")
	assert.True(t, ok)
	assert.NoError(t, err)

	// error - does not exist
	ok, err = New("/nope", 0666).Search("test")
	assert.False(t, ok)
	assert.EqualError(t, err, `cannot search note "nope" - does not exist`)
}

func TestUpdate(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Update("Test.\n")
	test.AssertFile(t, note.Orig, "Test.\n")
	assert.NoError(t, err)

	// error - does not exist
	err = New("/nope", 0666).Update("Test.\n")
	assert.EqualError(t, err, `cannot update note "nope" - does not exist`)
}
