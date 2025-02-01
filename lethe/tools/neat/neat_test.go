package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tTest.\n")
	assert.Equal(t, "Test.\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("\t.TEST\n")
	assert.Equal(t, ".test", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tTEST 123!!!\n")
	assert.Equal(t, "test-123", name)
}

func TestPath(t *testing.T) {
	// success
	path := Path("\t/dire/././name.extn\n")
	assert.Equal(t, "/dire/name.extn", path)
}
