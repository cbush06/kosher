package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cbush06/kosher/common"

	"github.com/cbush06/kosher/fs"
	"github.com/spf13/afero"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	fs.MockFs = afero.NewMemMapFs()
	common.BuildTestProject(fs.MockFs)
	Init()
	libPath := filepath.Join(filepath.Dir(""), "libs")
	assert.Containsf(t, os.Getenv("PATH"), libPath, "expected OS PATH env variable to contain [%s], but did not", libPath)
}
