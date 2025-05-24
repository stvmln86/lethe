package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	// success
	base := Base("/dire/name.extn")
	assert.Equal(t, "name.extn", base)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - real extension
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)

	// success - empty extension
	extn = Extn("/dire/name.")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("/dire/name")
	assert.Equal(t, ".", extn)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", dest)
}

func TestName(t *testing.T) {
	// success - real extension
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)

	// success - empty extension
	name = Name("/dire/name.")
	assert.Equal(t, "name", name)

	// success - no extension
	name = Name("/dire/name")
	assert.Equal(t, "name", name)
}

func TestRename(t *testing.T) {
	// success - dire
	dest := Rename("/dire/name.extn", "dire", "/test")
	assert.Equal(t, "/test/name.extn", dest)

	// success - name
	dest = Rename("/dire/name.extn", "name", "test")
	assert.Equal(t, "/dire/test.extn", dest)

	// success - extn
	dest = Rename("/dire/name.extn", "extn", ".test")
	assert.Equal(t, "/dire/name.test", dest)

	// panic - invalid mode
	defer func() {
		text := recover().(string)
		assert.Equal(t, "invalid rename mode", text)
	}()
	Rename("/dire/name.extn", "nope", "test")
}
