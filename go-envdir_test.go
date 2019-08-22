package main

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestFullCopy(t *testing.T) {

	testVars := []string{"AAA=aaa", "BBB123=bbb123"}

	envVars, err := readVars("testdata")
	assert.NilError(t, err)
	assert.Equal(t, true, reflect.DeepEqual(testVars, envVars))
}
