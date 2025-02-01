package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dire := test.TempDire(t)
	dest := filepath.Join(dire, "create.extn")

	// success
	err := Create(dest, "Test.\n", 0666)
	test.AssertFile(t, dest, "Test.\n")
	assert.NoError(t, err)

	// failure - already exists
	err = Create(dest, "Test.\n", 0666)
	test.AssertErr(t, err, "cannot create file %q - already exists")
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.NoError(t, err)

	// failure - does not exist
	err = Delete(orig)
	test.AssertErr(t, err, "cannot delete file %q - does not exist")
}

func TestExists(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope")
	assert.False(t, ok)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)

	// failure - does not exist
	body, err = Read("/nope")
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read file %q - does not exist")
}

func TestReextn(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")
	dest := filepath.Join(filepath.Dir(orig), "alpha.test")

	// success
	err := Reextn(orig, ".test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Reextn("/nope", ".test")
	test.AssertErr(t, err, "cannot rename file %q - does not exist")
}

func TestRename(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")
	dest := filepath.Join(filepath.Dir(orig), "test.extn")

	// success
	err := Rename(orig, "test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Rename("/nope", "test")
	test.AssertErr(t, err, "cannot rename file %q - does not exist")
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	err := Update(orig, "Test.\n", 0666)
	test.AssertFile(t, orig, "Test.\n")
	assert.NoError(t, err)

	// failure - does not exist
	err = Update("/nope", "Test.\n", 0666)
	test.AssertErr(t, err, "cannot update file %q - does not exist")
}
