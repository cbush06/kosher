package report

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"time"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/resources/reporttemplates"
	"github.com/spf13/afero"
)

const errMsg = "Error encountered while generating HTML report: %s"

var leadingWhitespace = regexp.MustCompile(`(?m)^(?:\s+)(.*)`)

// cukeComment is any single-line comment.
type cukeComment struct {
	Value string `json:"value"`
	Line  int    `json:"line"`
}

// cukeDocstring is a docstring (multi-line comment).
type cukeDocstring struct {
	Value       string `json:"value"`
	ContentType string `json:"content_type"`
	Line        int    `json:"line"`
}

// cuckeTag is any tag added to a feature or scenario.
type cukeTag struct {
	Name string `json:"name"`
	Line int    `json:"line"`
}

// cukeResult is the result of executing a step.
type cukeResult struct {
	Status   string `json:"status"`
	Error    string `json:"error_message,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

func (r *cukeResult) GetDurationInSeconds() string {
	if r.Duration != nil {
		return fmt.Sprintf("%0.2fs", time.Duration(*r.Duration).Seconds())
	}
	return ""
}

// cuckeMatch is the method matched to a step.
type cukeMatch struct {
	Location string `json:"location"`
}

// cukeEmbedding is any file or binary content attached to a step.
type cukeEmbedding struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

// cukeStep is a single step in a cukeElement.
type cukeStep struct {
	Keyword    string          `json:"keyword"`
	Name       string          `json:"name"`
	Line       int             `json:"line"`
	Docstring  *cukeDocstring  `json:"doc_string,omitempty"`
	Match      cukeMatch       `json:"match"`
	Result     cukeResult      `json:"result"`
	Embeddings []cukeEmbedding `json:"embeddings,omitempty"`
}

// cukeElement represents any block nested within a Feature:
//		* Background
//		* Scenario
//		* Scneario Outline
type cukeElement struct {
	ID           string     `json:"id"`
	Keyword      string     `json:"keyword"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Line         int        `json:"line"`
	Type         string     `json:"type"`
	Tags         []cukeTag  `json:"tags,omitempty"`
	Steps        []cukeStep `json:"steps,omitempty"`
	StepsPassed  int        `json:"-"`
	StepsFailed  int        `json:"-"`
	StepsPending int        `json:"-"`
	StepsSkipped int        `json:"-"`
}

func (e *cukeElement) GetTrimmedDescription() string {
	return leadingWhitespace.ReplaceAllString(e.Description, "$1")
}

// cukeFeature is a single feature in JSONReport.
type cukeFeature struct {
	URI             string        `json:"uri"`
	ID              string        `json:"id"`
	Keyword         string        `json:"keyword"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Line            int           `json:"line"`
	Comments        []cukeComment `json:"comments,omitempty"`
	Tags            []cukeTag     `json:"tags,omitempty"`
	Elements        []cukeElement `json:"elements,omitempty"`
	ElementsPassed  int           `json:"-"`
	ElementsFailed  int           `json:"-"`
	ElementsPending int           `json:"-"`
	StepsPassed     int           `json:"-"`
	StepsFailed     int           `json:"-"`
	StepsPending    int           `json:"-"`
	StepsSkipped    int           `json:"-"`
}

// GetTrimmedDescription returns the features description after removing leading whitespace from each line.
func (f *cukeFeature) GetTrimmedDescription() string {
	return leadingWhitespace.ReplaceAllString(f.Description, "$1")
}

// HTMLReport holds the jsonResults of a test execution
type HTMLReport struct {
	fileSystem      *fs.Fs
	settings        *config.Settings
	jsonResults     []byte
	Features        []cukeFeature
	ProjectName     string
	AppVersion      string
	Environment     string
	Browser         string
	Platform        string
	Timestamp       string
	ElementsPassed  int
	ElementsFailed  int
	ElementsPending int
	StepsPassed     int
	StepsFailed     int
	StepsPending    int
	StepsSkipped    int
	TotalElements   int
	TotalSteps      int
}

func newHTMLReport(s *config.Settings, f *fs.Fs) *HTMLReport {
	return &HTMLReport{
		fileSystem:      f,
		settings:        s,
		ProjectName:     s.Settings.GetString("projectName"),
		AppVersion:      s.Settings.GetString("appVersion"),
		Environment:     s.Settings.GetString("environment"),
		Browser:         "",
		Platform:        s.Settings.GetString("platform"),
		Timestamp:       "",
		ElementsPassed:  0,
		ElementsFailed:  0,
		ElementsPending: 0,
		StepsPassed:     0,
		StepsFailed:     0,
		StepsPending:    0,
		StepsSkipped:    0,
		TotalElements:   0,
		TotalSteps:      0,
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
	)

	reportFormat := r.settings.Settings.GetString("reportFormat")

	switch reportFormat {
	case "html", "bootstrap":
		templ, _ = template.New("Bootstrap").Parse(reporttemplates.GetBootstrapTemplate())
	default:
		return fmt.Errorf(errMsg, "attempt made to generate HTML report with unrecognized template")
	}

	if err = r.unmarshallJSON(); err != nil {
		return err
	}

	filePath, _ := r.fileSystem.ResultsDir.RealPath("results.html")
	if fileHandle, err = r.fileSystem.ResultsDir.OpenFile("results.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777); err != nil {
		return fmt.Errorf(fmt.Sprintf(errMsg, "failed to open results file [%s]: %s"), filePath, err)
	}

	if err = templ.Execute(fileHandle, r); err != nil {
		return fmt.Errorf(fmt.Sprintf(errMsg, "failed to generate report file [%s]: %s"), filePath, err)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf(fmt.Sprintf(errMsg, "failed to close report file [%s]: %s"), filePath, err)
	}
	return nil
}

func (r *HTMLReport) unmarshallJSON() error {
	if err := json.Unmarshal(r.jsonResults, &r.Features); err != nil {
		return fmt.Errorf(fmt.Sprintf(errMsg, "failed to parse JSON results of test execution: %s"), err)
	}

	for f := 0; f < len(r.Features); f++ {
		feature := &r.Features[f]

		for e := 0; e < len(feature.Elements); e++ {
			element := &feature.Elements[e]

			for s := 0; s < len(element.Steps); s++ {
				step := &element.Steps[s]

				switch step.Result.Status {
				case "passed":
					element.StepsPassed++
				case "failed":
					element.StepsFailed++
				case "skipped":
					element.StepsSkipped++
				case "undefined":
					element.StepsPending++
				}

				r.TotalSteps++
			}

			feature.StepsPassed += element.StepsPassed
			feature.StepsFailed += element.StepsFailed
			feature.StepsSkipped += element.StepsSkipped
			feature.StepsPending += element.StepsPending

			r.StepsPassed += element.StepsPassed
			r.StepsFailed += element.StepsFailed
			r.StepsSkipped += element.StepsSkipped
			r.StepsPending += element.StepsPending

			if element.StepsFailed > 0 {
				feature.ElementsFailed++
			} else if element.StepsPending > 0 {
				feature.ElementsPending++
			} else {
				feature.ElementsPassed++
			}

			r.TotalElements++
		}

		r.ElementsPassed += feature.ElementsPassed
		r.ElementsFailed += feature.ElementsFailed
		r.ElementsPending += feature.ElementsPending
	}

	return nil
}
