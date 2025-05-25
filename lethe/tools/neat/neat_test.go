package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("\t.EXTN\n")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123 !!!\n")
	assert.Equal(t, "name-123", name)
}

func TestPath(t *testing.T) {
	// success
	path := Path("\t/././path\n")
	assert.Equal(t, "/path", path)
}

func TestSubstring(t *testing.T) {
	// success - true
	ok := Substring("TEST", "t")
	assert.True(t, ok)

	// success - false
	ok = Substring("TEST", "nope")
	assert.False(t, ok)
}
