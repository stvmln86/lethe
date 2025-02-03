package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/lethe/lethe/tools/test"
)

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("TEST", "\tTest.\n")
	os.Setenv("BLANK", "\n")

	// success
	data, err := Env("TEST")
	assert.Equal(t, "Test.", data)
	assert.NoError(t, err)

	// failure - does not exist
	data, err = Env("NOPE")
	assert.Empty(t, data)
	test.AssertErr(t, err, "cannot read variable %q - does not exist")

	// failure - is empty
	data, err = Env("BLANK")
	assert.Empty(t, data)
	test.AssertErr(t, err, "cannot read variable %q - is empty")
}

func TestParse(t *testing.T) {
	// success - real argument
	amap, err := Parse([]string{"PARAMETER"}, []string{"ARGUMENT"})
	assert.Equal(t, map[string]string{"PARAMETER": "ARGUMENT"}, amap)
	assert.NoError(t, err)

	// success - default argument
	amap, err = Parse([]string{"PARAMETER:default"}, nil)
	assert.Equal(t, map[string]string{"PARAMETER": "default"}, amap)
	assert.NoError(t, err)

	// error - missing
	amap, err = Parse([]string{"PARAMETER"}, nil)
	assert.Empty(t, amap)
	test.AssertErr(t, err, "cannot parse arguments - %q missing")
}

func TestSplit(t *testing.T) {
	// success - empty
	name, argus := Split(nil)
	assert.Equal(t, "help", name)
	assert.Empty(t, argus)

	// success - with command
	name, argus = Split([]string{"test"})
	assert.Equal(t, "test", name)
	assert.Empty(t, argus)

	// success - with command and arguments
	name, argus = Split([]string{"test", "test"})
	assert.Equal(t, "test", name)
	assert.Equal(t, []string{"test"}, argus)
}
