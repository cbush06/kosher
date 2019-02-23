package report

import (
	"fmt"
	"html/template"
	"os"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/resources/reporttemplates"
	"github.com/spf13/afero"
)

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

// cukeFeature is a single feature in JSONReport.
type cukeFeature struct {
	URI            string        `json:"uri"`
	ID             string        `json:"id"`
	Keyword        string        `json:"keyword"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Line           int           `json:"line"`
	Comments       []cukeComment `json:"comments,omitempty"`
	Tags           []cukeTag     `json:"tags,omitempty"`
	Elements       []cukeElement `json:"elements,omitempty"`
	ElementsPassed int           `json:"-"`
	ElementsFailed int           `json:"-"`
}

// HTMLReport holds the jsonResults of a test execution
type HTMLReport struct {
	fileSystem     *fs.Fs
	settings       *config.Settings
	jsonResults    []byte
	Features       []cukeFeature
	ProjectName    string
	AppVersion     string
	Environment    string
	Browser        string
	Platform       string
	Timestamp      string
	FeaturesPassed int
	FeaturesFailed int
	StepsPassed    int
	StepsFailed    int
	StepsPending   int
	StepsSkipped   int
}

func newHTMLReport(s *config.Settings, f *fs.Fs) *HTMLReport {
	return &HTMLReport{
		fileSystem:     f,
		settings:       s,
		ProjectName:    s.Settings.GetString("projectName"),
		AppVersion:     s.Settings.GetString("appVersion"),
		Environment:    s.Settings.GetString("environment"),
		Browser:        "",
		Platform:       s.Settings.GetString("platform"),
		Timestamp:      "",
		FeaturesPassed: 0,
		FeaturesFailed: 0,
		StepsPassed:    0,
		StepsFailed:    0,
		StepsPending:   0,
		StepsSkipped:   0,
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
	}

	filePath, _ := r.fileSystem.ResultsDir.RealPath("results.html")
	if fileHandle, err = r.fileSystem.ResultsDir.OpenFile("results.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777); err != nil {
		return fmt.Errorf("Error encountered opening results file [%s]: %s", filePath, err)
	}
	if err = templ.Execute(fileHandle, r); err != nil {
		return fmt.Errorf("Error encountered generating report file [%s]: %s", filePath, err)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf("Error encountered closing report file [%s]: %s", filePath, err)
	}
	return nil
}
