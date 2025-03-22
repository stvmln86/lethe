package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtn(t *testing.T) {
	// success - with extension
	extn := Extn("/dire/slug.extn")
	assert.Equal(t, "extn", extn)

	// success - with dot only
	extn = Extn("/dire/slug.")
	assert.Equal(t, "", extn)

	// success - without extension
	extn = Extn("/dire/slug")
	assert.Equal(t, "", extn)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "slug", "extn")
	assert.Equal(t, "/dire/slug.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	okay := Match("/dire/slug.extn", "SLU")
	assert.True(t, okay)

	// success - false
	okay = Match("/dire/slug.extn", "NOPE")
	assert.False(t, okay)
}

func TestReextn(t *testing.T) {
	// success
	dest := Reextn("/dire/slug.extn", "test")
	assert.Equal(t, "/dire/slug.test", dest)
}

func TestReslug(t *testing.T) {
	// success
	dest := Reslug("/dire/slug.extn", "test")
	assert.Equal(t, "/dire/test.extn", dest)
}

func TestSlug(t *testing.T) {
	// success - with extension
	slug := Slug("/dire/slug.extn")
	assert.Equal(t, "slug", slug)

	// success - with dot only
	slug = Slug("/dire/slug.")
	assert.Equal(t, "slug", slug)

	// success - without extension
	slug = Slug("/dire/slug")
	assert.Equal(t, "slug", slug)
}
