package common

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
