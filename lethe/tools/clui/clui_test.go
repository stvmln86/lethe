package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("NAME", "Data.\n")
	os.Setenv("BLANK", "\n")

	// success
	data, err := Env("NAME")
	assert.Equal(t, "Data.", data)
	assert.NoError(t, err)

	// error - is blank
	data, err = Env("BLANK")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variable "BLANK" is blank`)

	// error - is not set
	data, err = Env("NOPE")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variable "NOPE" is not set`)
}

func TestParse(t *testing.T) {
	// success - with argument
	pairs, err := Parse([]string{"NAME"}, []string{"argument"})
	assert.Equal(t, map[string]string{"NAME": "argument"}, pairs)
	assert.NoError(t, err)

	// success - with default
	pairs, err = Parse([]string{"NAME:default"}, nil)
	assert.Equal(t, map[string]string{"NAME": "default"}, pairs)
	assert.NoError(t, err)

	// error - not provided
	pairs, err = Parse([]string{"NAME"}, nil)
	assert.Empty(t, pairs)
	assert.EqualError(t, err, `argument "NAME" was not provided`)
}

func TestSplit(t *testing.T) {
	// success - no arguments
	name, elems := Split(nil)
	assert.Empty(t, name)
	assert.Empty(t, elems)

	// success - one argument
	name, elems = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Empty(t, elems)

	// success - multiple arguments
	name, elems = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, elems)
}
