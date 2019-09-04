package report

import (
	"fmt"
	"html/template"
	"os"
	"runtime"
	"time"

	"github.com/cbush06/kosher/config"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/resources/reporttemplates"
	"github.com/spf13/afero"
)

const htmlErrMsg = "Error encountered while generating HTML report: %s"

// HTMLReport holds the jsonResults of a test execution
type HTMLReport struct {
	CucumberReport
	jsonResults []byte
}

func newHTMLReport(s *config.Settings) *HTMLReport {
	return &HTMLReport{
		CucumberReport: NewCucumberReport(s),
	}
}

// Write implements the io.Writer interface for the JSONReport struct. This
// allows JSONReport to be the Output stream for GoDog.
func (r *HTMLReport) Write(p []byte) (int, error) {
	r.jsonResults = append(r.jsonResults, p...)
	return len(p), nil
}

// Process marshalls the `jsonResults` into structs and analyzes the results
// to generate metrics such as `StepsPassed`, `StepsFailed`, etc.
func (r *HTMLReport) Process() error {
	var (
		templ      *template.Template
		fileHandle afero.File
		err        error
		writeLen   int
	)

	const (
		resultsHTMLFile = common.ResultsHTMLFile
		resultsJSONFile = common.ResultsJSONFile
	)

	reportFormat := r.settings.Settings.GetString("reportFormat")
	r.Timestamp = time.Now().Format(time.RFC850)
	r.OS = fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH)

	// build template and parse/unmarshall JSON results
	switch reportFormat {
	case "html", "bootstrap":
		templ, _ = template.New("Bootstrap").Parse(reporttemplates.GetBootstrapTemplate())
	case "simple":
		templ, _ = template.New("Simple").Parse(reporttemplates.GetSimpleTemplate())
	default:
		return fmt.Errorf(htmlErrMsg, "attempt made to generate HTML report with unrecognized template")
	}

	if err = r.UnmarshallJSON(r.jsonResults); err != nil {
		return err
	}

	// write HTML report file
	filePath, _ := r.settings.FileSystem.ResultsDir.RealPath(resultsHTMLFile)
	if fileHandle, err = r.settings.FileSystem.ResultsDir.OpenFile(resultsHTMLFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to open results file [%s]: %s"), filePath, err)
	}
	if err = templ.Execute(fileHandle, r); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to generate report file [%s]: %s"), filePath, err)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to close report file [%s]: %s"), filePath, err)
	}

	// write JSON results file
	filePath, _ = r.settings.FileSystem.ResultsDir.RealPath(resultsJSONFile)
	if fileHandle, err = r.settings.FileSystem.ResultsDir.OpenFile(resultsJSONFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to open results file [%s]: %s"), filePath, err)
	}
	if writeLen, err = fileHandle.Write(r.jsonResults); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to write report file [%s]: %s"), filePath, err)
	} else if writeLen < len(r.jsonResults) {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to write all bytes of report file [%s]"), filePath)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf(fmt.Sprintf(htmlErrMsg, "failed to close report file [%s]: %s"), filePath, err)
	}

	return nil
}
