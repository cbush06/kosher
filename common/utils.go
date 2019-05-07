package common

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/spf13/afero"
)

// StringSliceContains returns true if the SORTED slice `s` contains `v`.
func StringSliceContains(s []string, v string) bool {
	idx := sort.SearchStrings(s, v)
	return idx < len(s) && s[idx] == v
}

// StringSliceContainsFold returns true if the SORTED slice `s` contains `v`, ignoring case.
func StringSliceContainsFold(s []string, v string) bool {
	vUpper := strings.ToUpper(v)
	sUpper := s[:]
	for i, next := range sUpper {
		sUpper[i] = strings.ToUpper(next)
	}
	idx := sort.SearchStrings(sUpper, vUpper)
	return idx < len(sUpper) && sUpper[idx] == vUpper
}

// BuildTestProject is intended for using during unit testing to build an in-memory representation of the project
func BuildTestProject(fs afero.Fs) error {
	workingDir, _ := os.Getwd()
	fs.MkdirAll(workingDir, 0777)
	fs.MkdirAll(filepath.Join(workingDir, ConfigDir), 0777)
	fs.MkdirAll(filepath.Join(workingDir, FeaturesDir), 0777)
	fs.MkdirAll(filepath.Join(workingDir, ResultsDir), 0777)
	fs.MkdirAll(filepath.Join(workingDir, MacrosDir), 0777)
	fs.MkdirAll(filepath.Join(workingDir, ResourcesDir), 0777)
	afero.WriteReader(fs, filepath.Join(workingDir, ConfigDir, EnvironmentsFile), bytes.NewBufferString(`{"test":""}`))
	afero.WriteReader(fs, filepath.Join(workingDir, ConfigDir, PagesFile), bytes.NewBufferString(`{}`))
	afero.WriteReader(fs, filepath.Join(workingDir, ConfigDir, SelectorsFile), bytes.NewBufferString(`{}`))
	afero.WriteReader(fs, filepath.Join(workingDir, ConfigDir, SettingsFile), bytes.NewBufferString(`{}`))
	return nil
}

// CaptureStdout executes a task and returns any text it wrote to Stdout
func CaptureStdout(task func()) string {
	reader, writer, _ := os.Pipe()

	stdout := os.Stdout
	defer func() {
		reader.Close()
		os.Stdout = stdout
		log.SetOutput(stdout)
	}()
	os.Stdout = writer
	log.SetOutput(writer)

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
	task()
	writer.Close()
	return <-capture
}
