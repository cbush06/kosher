package report

import (
	"github.com/cbush06/kosher/fs"
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
	ID          string     `json:"id"`
	Keyword     string     `json:"keyword"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Line        int        `json:"line"`
	Type        string     `json:"type"`
	Tags        []cukeTag  `json:"tags,omitempty"`
	Steps       []cukeStep `json:"steps,omitempty"`
}

// cukeFeature is a single feature in JSONReport.
type cukeFeature struct {
	URI         string        `json:"uri"`
	ID          string        `json:"id"`
	Keyword     string        `json:"keyword"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Line        int           `json:"line"`
	Comments    []cukeComment `json:"comments,omitempty"`
	Tags        []cukeTag     `json:"tags,omitempty"`
	Elements    []cukeElement `json:"elements,omitempty"`
}

// JSONReport holds the jsonResults of a test execution
type JSONReport struct {
	jsonResults    []byte
	features       []cukeFeature
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

func NewJsonReport(fs *fs.Fs, s *config.S)

// Write implements the io.Writer interface for the JSONReport struct. This
// allows JSONReport to be the Output stream for GoDog.
func (r *JSONReport) Write(p []byte) (int, error) {
	r.jsonResults = append(r.jsonResults, p...)
	return len(p), nil
}

// Process marshalls the `jsonResults` into structs and analyzes the results
// to generate metrics such as `StepsPassed`, `StepsFailed`, etc.
func (r *JSONReport) Process() error {

	return nil
}
