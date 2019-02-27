package report

import (
	"log"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
)

// Report is a struct that receives data from GoDog via the `Write` method. Depending on the
// report type, it will either pass through the data to its final destination or cache it,
// format it, and then write it out when the Process() method is called.
type Report interface {
	Write(p []byte) (int, error)
	Process() error
}

// NewReport is a factory method that creates the appropriate `Report` object based on
// the system's settings.
func NewReport(s *config.Settings, f *fs.Fs) Report {
	reportFormat := s.Settings.GetString("reportFormat")

	switch reportFormat {
	case "html", "bootstrap", "simple":
		return newHTMLReport(s, f)
	case "pretty", "progress":
		return newStdOutWriteThroughReport()
	case "junit":
		if report, err := newFileWriteThroughReport("results.xml", f); err != nil {
			log.Fatalf("Error encountered while building JUnit report: %s\n", err)
		} else {
			return report
		}
	case "cucumber":
		if report, err := newFileWriteThroughReport("results.json", f); err != nil {
			log.Fatalf("Error encountered while building Cucumber report: %s\n", err)
		} else {
			return report
		}
	}

	return nil
}
