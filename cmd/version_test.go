package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/cbush06/kosher/common"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	cmdVersion := buildVersionCommand()

	reader, writer, _ := os.Pipe()
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()
	os.Stdout = writer

	capture := make(chan string)

	wait := new(sync.WaitGroup)
	wait.Add(1)
	go func() {
		buf := bytes.NewBuffer([]byte{})
		wait.Done()
		io.Copy(buf, reader)
		capture <- buf.String()
	}()
	wait.Wait()
	cmdVersion.command.Execute()
	writer.Close()
	assert.Equal(t, common.CurrentVersion.Version(), strings.TrimSpace(<-capture))
	reader.Close()
}
