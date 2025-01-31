package path

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func TestBase(t *testing.T) {
	// success
	base := Base("/dire/name.extn")
	assert.Equal(t, "name.extn", base)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.TempDire(t)

	// success
	origs := Glob(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
		filepath.Join(dire, "charlie.extn"),
	}, origs)
}

func TestJoin(t *testing.T) {
	// success
	orig := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", orig)
}

func TestName(t *testing.T) {
	// success
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)
}

func TestReextn(t *testing.T) {
	// success
	dest := Reextn("/dire/name.extn", ".test")
	assert.Equal(t, "/dire/name.test", dest)
}

func TestRename(t *testing.T) {
	// success
	dest := Rename("/dire/name.extn", "test")
	assert.Equal(t, "/dire/test.extn", dest)
}
