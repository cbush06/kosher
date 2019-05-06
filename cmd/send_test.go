package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendCommand(t *testing.T) {
	cmdSend := buildSendCommand()
	result := cmdSend.command.Execute()
	assert.Nilf(t, result, "expected send command to succeed, but did not: %s", result)
}
