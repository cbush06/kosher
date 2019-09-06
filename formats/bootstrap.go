package formats

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"text/template"
	"time"

	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/resources/reporttemplates"
	"github.com/spf13/afero"
)

// NewBootstrapFormatterFunc builds a new FormatterFunc to be used when registering BootstrapFormat with godog
func NewBootstrapFormatterFunc(s *config.Settings) godog.FormatterFunc {
	return func(suite string, out io.Writer) godog.Formatter {
		return &BootstrapFormat{
			CucumberFormatter: NewCucumberFormatter(s),
		}
	}
}

// BootstrapFormat ...
type BootstrapFormat struct {
	CucumberFormatter
}

// Feature ...
func (b *BootstrapFormat) Feature(f *gherkin.Feature, p string, src []byte) {
	b.CucumberFormatter.Feature(f, p, src)
}

// Node ...
func (b *BootstrapFormat) Node(n interface{}) {
	b.CucumberFormatter.Node(n)
	switch n.(type) {
	case gherkin.Scenario, gherkin.ScenarioOutline:
	}
}

// Defined ...
func (b *BootstrapFormat) Defined(s *gherkin.Step, sd *godog.StepDef) {
	b.CucumberFormatter.Defined(s, sd)
}

// Failed ...
func (b *BootstrapFormat) Failed(s *gherkin.Step, sd *godog.StepDef, err error) {
	b.CucumberFormatter.Failed(s, sd, err)
}

// Passed ...
func (b *BootstrapFormat) Passed(s *gherkin.Step, sd *godog.StepDef) {
	b.CucumberFormatter.Passed(s, sd)
}

// Skipped ...
func (b *BootstrapFormat) Skipped(s *gherkin.Step, sd *godog.StepDef) {
	b.CucumberFormatter.Skipped(s, sd)
	b.StepsSkipped++
}

// Undefined ...
func (b *BootstrapFormat) Undefined(s *gherkin.Step, sd *godog.StepDef) {
	b.CucumberFormatter.Undefined(s, sd)
}

// Pending ...
func (b *BootstrapFormat) Pending(s *gherkin.Step, sd *godog.StepDef) {
	b.CucumberFormatter.Pending(s, sd)
	b.StepsPending++
}

// Summary ...
func (b *BootstrapFormat) Summary() {
	b.CucumberFormatter.Summary()

	var (
		templ      *template.Template
		fileHandle afero.File
		err        error
	)

	b.Timestamp = time.Now().Format(time.RFC850)
	b.OS = fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH)

	// build template and parse/unmarshall JSON results
	templ, _ = template.New("Bootstrap").Parse(reporttemplates.GetBootstrapTemplate())

	// write HTML report file
	filePath, _ := b.Settings.FileSystem.ResultsDir.RealPath(common.ResultsBootstrapFile)
	if fileHandle, err = b.Settings.FileSystem.ResultsDir.OpenFile(common.ResultsBootstrapFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
		log.Printf(fmt.Sprintf(htmlErrMsg, "failed to open results file [%s]: %s"), filePath, err)
		return
	}
	if err = templ.Execute(fileHandle, b); err != nil {
		log.Printf(fmt.Sprintf(htmlErrMsg, "failed to generate report file [%s]: %s"), filePath, err)
		return
	}
	if err = fileHandle.Close(); err != nil {
		log.Printf(fmt.Sprintf(htmlErrMsg, "failed to close report file [%s]: %s"), filePath, err)
		return
	}
}
