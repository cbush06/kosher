package report

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/cbush06/kosher/config"
)

var leadingWhitespace = regexp.MustCompile(`(?m)^(?:\s*)(.*\S)(?:\s*)$`)

// CukeComment is any single-line comment.
type CukeComment struct {
	Value string `json:"value"`
	Line  int    `json:"line"`
}

// CukeDocstring is a docstring (multi-line comment).
type CukeDocstring struct {
	Value       string `json:"value"`
	ContentType string `json:"content_type"`
	Line        int    `json:"line"`
}

// CukeTag is any tag added to a feature or scenario.
type CukeTag struct {
	Name string `json:"name"`
	Line int    `json:"line"`
}

// CukeResult is the result of executing a step.
type CukeResult struct {
	Status   string `json:"status"`
	Error    string `json:"error_message,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

// GetDurationInSeconds convers the millisecond representation of duration to seconds (accurate to 2 decimal places) and returns the result.
func (r *CukeResult) GetDurationInSeconds() string {
	if r.Duration != nil {
		return fmt.Sprintf("%0.2fs", time.Duration(*r.Duration).Seconds())
	}
	return ""
}

// CukeMatch is the method matched to a step.
type CukeMatch struct {
	Location string `json:"location"`
}

// CukeEmbedding is any file or binary content attached to a step.
type CukeEmbedding struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

// CukeStep is a single step in a cukeElement.
type CukeStep struct {
	Keyword    string              `json:"keyword"`
	Name       string              `json:"name"`
	Line       int                 `json:"line"`
	Docstring  *CukeDocstring      `json:"doc_string,omitempty"`
	Match      CukeMatch           `json:"match"`
	Result     CukeResult          `json:"result"`
	Embeddings []CukeEmbedding     `json:"embeddings,omitempty"`
	DataTable  []*CukeDataTableRow `json:"rows,omitempty"`
}

// GetTrimmedKeyword removes leading and trailing whitespace from the Step's keyword.
func (s *CukeStep) GetTrimmedKeyword() string {
	return leadingWhitespace.ReplaceAllString(s.Keyword, "$1")
}

// GetTrimmedName removes leading and trailing whitespace from the Step's name.
func (s *CukeStep) GetTrimmedName() string {
	return leadingWhitespace.ReplaceAllString(s.Name, "$1")
}

// CukeDataTableRow represents a row in a DataTable owned by a step
type CukeDataTableRow struct {
	Cells []string `json:"cells"`
}

// CukeElement represents any block nested within a Feature:
//		* Background
//		* Scenario
//		* Scneario Outline
type CukeElement struct {
	ID           string     `json:"id"`
	Keyword      string     `json:"keyword"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Line         int        `json:"line"`
	Type         string     `json:"type"`
	Tags         []CukeTag  `json:"tags,omitempty"`
	Steps        []CukeStep `json:"steps,omitempty"`
	StepsPassed  int        `json:"-"`
	StepsFailed  int        `json:"-"`
	StepsPending int        `json:"-"`
	StepsSkipped int        `json:"-"`
}

// GetTrimmedKeyword removes leading and trailing whitespace from the Scenario's keyword.
func (e *CukeElement) GetTrimmedKeyword() string {
	return leadingWhitespace.ReplaceAllString(e.Keyword, "$1")
}

// GetTrimmedDescription removes leading and trailing whitespace from the Scenario's description and returns the result.
func (e *CukeElement) GetTrimmedDescription() string {
	return leadingWhitespace.ReplaceAllString(e.Description, "$1")
}

// CukeFeature is a single feature in JSONReport.
type CukeFeature struct {
	URI             string        `json:"uri"`
	ID              string        `json:"id"`
	Keyword         string        `json:"keyword"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Line            int           `json:"line"`
	Comments        []CukeComment `json:"comments,omitempty"`
	Tags            []CukeTag     `json:"tags,omitempty"`
	Elements        []CukeElement `json:"elements,omitempty"`
	ElementsPassed  int           `json:"-"`
	ElementsFailed  int           `json:"-"`
	ElementsPending int           `json:"-"`
	StepsPassed     int           `json:"-"`
	StepsFailed     int           `json:"-"`
	StepsPending    int           `json:"-"`
	StepsSkipped    int           `json:"-"`
}

// GetTrimmedDescription returns the features description after removing leading and trailing whitespace from each line.
func (f *CukeFeature) GetTrimmedDescription() string {
	return leadingWhitespace.ReplaceAllString(f.Description, "$1")
}

// CucumberReport holds the jsonResults of a test execution
type CucumberReport struct {
	settings        *config.Settings
	Features        []CukeFeature
	ProjectName     string
	AppVersion      string
	Environment     string
	Browser         string
	Platform        string
	RunTime         time.Duration
	OS              string
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

// NewCucumberReport creates a new CucumberReport struct and initializes it.
func NewCucumberReport(s *config.Settings) CucumberReport {
	return CucumberReport{
		settings:        s,
		ProjectName:     s.Settings.GetString("projectName"),
		AppVersion:      s.Settings.GetString("appVersion"),
		Environment:     s.Settings.GetString("environment"),
		Browser:         s.Settings.GetString("driver"),
		Platform:        s.Settings.GetString("platform"),
		RunTime:         time.Duration(0),
		OS:              "",
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

// UnmarshallJSON unmarshalls JSON content into an array of CukeFeatures.
func (r *CucumberReport) UnmarshallJSON(jsonResults []byte) error {
	if err := json.Unmarshal(jsonResults, &r.Features); err != nil {
		return fmt.Errorf(fmt.Sprintf(errMsg, "failed to parse JSON results of test execution: %s"), err)
	}

	for f := 0; f < len(r.Features); f++ {
		feature := &r.Features[f]

		for e := 0; e < len(feature.Elements); e++ {
			element := &feature.Elements[e]

			for s := 0; s < len(element.Steps); s++ {
				step := &element.Steps[s]

				if step.DataTable == nil {
					step.DataTable = make([]*CukeDataTableRow, 0)
				}

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

				if step.Result.Duration != nil {
					r.RunTime += time.Duration(*step.Result.Duration)
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
