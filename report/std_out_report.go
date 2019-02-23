package report

import (
	"fmt"
	"os"
)

// StdOutWriteThroughReport is an empty struct that serves to pass GoDog's writes directly to StdOut.
type StdOutWriteThroughReport struct{}

func newStdOutWriteThroughReport() *StdOutWriteThroughReport {
	return &StdOutWriteThroughReport{}
}

func (s *StdOutWriteThroughReport) Write(b []byte) (int, error) {
	fmt.Fprint(os.Stdout, string(b))
	return len(b), nil
}

// Process does nothing here. It simply implements the Report interface.
func (s *StdOutWriteThroughReport) Process() error {
	return nil
}
