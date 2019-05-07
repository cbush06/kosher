package cmd

import (
	"strings"
	"testing"

	"github.com/cbush06/kosher/common"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	cmdVersion := buildVersionCommand()
	output := common.CaptureStdout(func() { cmdVersion.command.Execute() })
	assert.Equal(t, common.CurrentVersion.Version(), strings.TrimSpace(output))
}
