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
	var (
		err        error
		fileHandle afero.File
	)

	if fileHandle, err = f.ResultsDir.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
		filePath, _ := f.ResultsDir.RealPath(fileName)
		return nil, fmt.Errorf("Unable to create results file [%s]: %s", filePath, err)
	}

	return &FileWriteThroughReport{
		fileHandle: fileHandle,
	}, nil
}

func (s *FileWriteThroughReport) Write(b []byte) (int, error) {
	return s.fileHandle.Write(b)
}

// Process does nothing here. It simply implements the Report interface.
func (s *FileWriteThroughReport) Process() error {
	return s.fileHandle.Close()
}
