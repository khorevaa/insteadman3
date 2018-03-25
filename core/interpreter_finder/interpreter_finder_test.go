package interpreterFinder

import (
	"../configurator"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFindAndCheckInterpreter(t *testing.T) {
	config := configurator.InsteadmanConfig{}
	finder := InterpreterFinder{Config: &config}

	interpreterPath := finder.Find()
	assert.NotNil(t, interpreterPath)

	config.InterpreterCommand = *interpreterPath
	version, e := finder.Check()

	assert.NoError(t, e)
	assert.Regexp(t, regexp.MustCompile("^\\d+.\\d+.\\d+"), version) // like "3.2.0"
}