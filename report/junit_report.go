package report

import (
	"fmt"
	"os"

	"github.com/cbush06/kosher/fs"

	"github.com/spf13/afero"
)

// FileWriteThroughReport is an empty struct that serves to pass GoDog's writes directly to a file.
type FileWriteThroughReport struct {
	fileHandle afero.File
}

func newFileWriteThroughReport(fileName string, f *fs.Fs) (*FileWriteThroughReport, error) {
	if fileHandle, err := f.ResultsDir.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0777); err != nil {
		return &FileWriteThroughReport{
			fileHandle: fileHandle,
		}, nil
	}
	filePath, _ := f.ResultsDir.RealPath("results.xml")
	return nil, fmt.Errorf("Unable to create results file [%s]", filePath)
}

func (s *FileWriteThroughReport) Write(b []byte) (int, error) {
	return s.fileHandle.Write(b)
}

// Process does nothing here. It simply implements the Report interface.
func (s *FileWriteThroughReport) Process() error {
	return s.fileHandle.Close()
}
