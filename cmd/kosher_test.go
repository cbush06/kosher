package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKosherCommand(t *testing.T) {
	cmdKosher := buildKosherCommand()
	result := cmdKosher.command.Execute()
	assert.Nilf(t, result, "expected kosher command to succeed, but did not: %s", result)
}
