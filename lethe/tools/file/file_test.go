package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope")
	assert.False(t, ok)
}

func TestList(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs := List(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, origs)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha", "test", 1)

	// success
	err := Rename(orig, "name", "test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Write(orig, "Test.\n", 0666)
	test.AssertFile(t, orig, "Test.\n")
	assert.NoError(t, err)
}
