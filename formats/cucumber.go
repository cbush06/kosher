package formats

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/afero"

	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/config"
)

// NewCucumberFormatterFunc builds a new FormatterFunc to be used when registering CucumberFormatter with godog
func NewCucumberFormatterFunc(s *config.Settings) godog.FormatterFunc {
	const resultsJSONFile = common.ResultsJSONFile

	return func(suite string, out io.Writer) godog.Formatter {
		var (
			filePath   string
			fileHandle afero.File
			err        error
		)

		// write JSON results file
		filePath, _ = s.FileSystem.ResultsDir.RealPath(resultsJSONFile)
		if fileHandle, err = s.FileSystem.ResultsDir.OpenFile(resultsJSONFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
			log.Fatalf(fmt.Sprintf(htmlErrMsg, "failed to open results file [%s]: %s"), filePath, err)
		}

		return &CucumberFormatter{
			Settings: s,
			out:      fileHandle,
		}
	}
}

// NewCucumberFormatter builds a new CucumberFormat struct
func NewCucumberFormatter(s *config.Settings) CucumberFormatter {
	return CucumberFormatter{
		CucumberReport: NewCucumberReport(s),
		Settings:       s,
	}
}

// CucumberFormatter builds out a data structure that maps to Cucumber JSON format.
type CucumberFormatter struct {
	CucumberReport
	Settings       *config.Settings
	owner          interface{}
	curStep        *CukeStep
	curElement     *CukeElement
	curFeature     *CukeFeature
	curOutline     CukeElement
	curRow         int
	curExampleName string
	curExampleTags []CukeTag
	startTime      time.Time
	out            io.Writer
}

type stepResult struct {
	Result string
	Error  error
}

func makeID(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}

// Feature records a new Feature for the Bootstrap report
func (b *CucumberFormatter) Feature(ft *gherkin.Feature, p string, c []byte) {
	newFeature := CukeFeature{
		Comments:    make([]CukeComment, 0),
		Description: ft.Description,
		Keyword:     ft.Keyword,
		Line:        ft.Location.Line,
		Name:        ft.Name,
		URI:         p,
		ID:          makeID(ft.Name),
	}

	for _, c := range ft.Comments {
		newFeature.Comments = append(newFeature.Comments, CukeComment{
			Line:  c.Location.Line,
			Value: c.Text,
		})
	}

	for _, t := range ft.Tags {
		newFeature.Tags = append(newFeature.Tags, CukeTag{
			Line: t.Location.Line,
			Name: t.Name,
		})
	}

	b.Features = append(b.Features, newFeature)
	b.curFeature = &b.Features[len(b.Features)-1]
	b.curElement = nil
}

// Node captures the Scenario or Outline that is the owner of the following steps
func (b *CucumberFormatter) Node(n interface{}) {
	switch t := n.(type) {

	// When the example definition is seen we just need track the id and
	// append the name associated with the example as part of the id.
	case *gherkin.Examples:

		b.curExampleName = makeID(t.Name)
		b.curRow = 2 // there can be more than one example set per outline so reset row count.
		// cucumber counts the header row as an example when creating the id.

		// store any example level tags in a  temp location.
		b.curExampleTags = make([]CukeTag, len(t.Tags))
		for idx, element := range t.Tags {
			b.curExampleTags[idx].Line = element.Location.Line
			b.curExampleTags[idx].Name = element.Name
		}

	// The outline node creates a placeholder and the actual element is added as each TableRow is processed.
	case *gherkin.ScenarioOutline:
		b.determineElementStatus()
		b.owner = t
		b.curOutline = CukeElement{}
		b.curOutline.Name = t.Name
		b.curOutline.Line = t.Location.Line
		b.curOutline.Description = t.Description
		b.curOutline.Keyword = t.Keyword
		b.curOutline.ID = b.curFeature.ID + ";" + makeID(t.Name)
		b.curOutline.Type = "scenario"
		b.curOutline.Tags = make([]CukeTag, len(t.Tags)+len(b.curFeature.Tags))

		// apply feature level tags
		if len(b.curOutline.Tags) > 0 {
			copy(b.curOutline.Tags, b.curFeature.Tags)

			// apply outline level tags.
			for idx, element := range t.Tags {
				b.curOutline.Tags[idx+len(b.curFeature.Tags)].Line = element.Location.Line
				b.curOutline.Tags[idx+len(b.curFeature.Tags)].Name = element.Name
			}
		}
		b.TotalElements++

	// This scenario adds the element to the output immediately.
	case *gherkin.Scenario:
		b.determineElementStatus()
		b.owner = t
		b.curFeature.Elements = append(b.curFeature.Elements, CukeElement{})
		b.curElement = &b.curFeature.Elements[len(b.curFeature.Elements)-1]

		b.curElement.Name = t.Name
		b.curElement.Line = t.Location.Line
		b.curElement.Description = t.Description
		b.curElement.Keyword = t.Keyword
		b.curElement.ID = b.curFeature.ID + ";" + makeID(t.Name)
		b.curElement.Type = "scenario"
		b.curElement.Tags = make([]CukeTag, len(t.Tags)+len(b.curFeature.Tags))

		if len(b.curElement.Tags) > 0 {
			// apply feature level tags
			copy(b.curElement.Tags, b.curFeature.Tags)

			// apply scenario level tags.
			for idx, element := range t.Tags {
				b.curElement.Tags[idx+len(b.curFeature.Tags)].Line = element.Location.Line
				b.curElement.Tags[idx+len(b.curFeature.Tags)].Name = element.Name
			}
		}
		b.TotalElements++

	// This is an outline scenario and the element is added to the output as
	// the TableRows are encountered.
	case *gherkin.TableRow:
		b.determineElementStatus()
		b.owner = t
		tmpElem := b.curOutline
		tmpElem.Line = t.Location.Line
		tmpElem.ID = tmpElem.ID + ";" + b.curExampleName + ";" + strconv.Itoa(b.curRow)
		b.curRow++
		b.curFeature.Elements = append(b.curFeature.Elements, tmpElem)
		b.curElement = &b.curFeature.Elements[len(b.curFeature.Elements)-1]

		// copy in example level tags.
		b.curElement.Tags = append(b.curElement.Tags, b.curExampleTags...)
	}
}

func (b *CucumberFormatter) determineElementStatus() {
	el := b.curElement
	ft := b.curFeature

	// skip the very first iteration
	if el == nil || ft == nil {
		return
	}

	if el.StepsFailed > 0 {
		ft.ElementsFailed++
		b.ElementsFailed++
	} else if el.StepsPending > 0 {
		ft.ElementsPending++
		b.ElementsPending++
	} else {
		ft.ElementsPassed++
		b.ElementsPassed++
	}
}

// Defined ...
func (b *CucumberFormatter) Defined(s *gherkin.Step, sd *godog.StepDef) {
	b.startTime = time.Now() // start timing the step
	b.curElement.Steps = append(b.curElement.Steps, CukeStep{})
	b.curStep = &b.curElement.Steps[len(b.curElement.Steps)-1]

	b.curStep.Name = s.Text
	b.curStep.Line = s.Location.Line
	b.curStep.Keyword = s.Keyword

	b.TotalSteps++

	if _, ok := s.Argument.(*gherkin.DocString); ok {
		b.curStep.Docstring = &CukeDocstring{}
		b.curStep.Docstring.ContentType = strings.TrimSpace(s.Argument.(*gherkin.DocString).ContentType)
		b.curStep.Docstring.Line = s.Argument.(*gherkin.DocString).Location.Line
		b.curStep.Docstring.Value = s.Argument.(*gherkin.DocString).Content
	}

	if _, ok := s.Argument.(*gherkin.DataTable); ok {
		dataTable := s.Argument.(*gherkin.DataTable)

		b.curStep.DataTable = make([]*CukeDataTableRow, len(dataTable.Rows))
		for i, row := range dataTable.Rows {
			cells := make([]string, len(row.Cells))
			for j, cell := range row.Cells {
				cells[j] = cell.Value
			}
			b.curStep.DataTable[i] = &CukeDataTableRow{Cells: cells}
		}
	}
}

func (b *CucumberFormatter) step(res *stepResult) {
	// determine if test case has finished
	switch t := b.owner.(type) {
	case *gherkin.TableRow:
		d := int(time.Now().Sub(b.startTime).Nanoseconds())
		b.curStep.Result.Duration = &d
		b.curStep.Line = t.Location.Line
		b.curStep.Result.Status = res.Result
		if res.Error != nil {
			b.curStep.Result.Error = res.Error.Error()
		}
	case *gherkin.Scenario, *gherkin.ScenarioOutline:
		d := int(time.Now().Sub(b.startTime).Nanoseconds())
		b.curStep.Result.Duration = &d
		b.curStep.Result.Status = res.Result
		if res.Error != nil {
			b.curStep.Result.Error = res.Error.Error()
		}
	}

	if b.curStep.DataTable == nil {
		b.curStep.DataTable = make([]*CukeDataTableRow, 0)
	}

	if b.curStep.Embeddings == nil {
		b.curStep.Embeddings = make([]*CukeEmbedding, 0)
	}
}

// Failed records a failed step
func (b *CucumberFormatter) Failed(s *gherkin.Step, sd *godog.StepDef, err error) {
	b.step(&stepResult{
		Error:  err,
		Result: "failed",
	})
	b.curElement.StepsFailed++
	b.curFeature.StepsFailed++
	b.StepsFailed++
}

// Passed records a passed step
func (b *CucumberFormatter) Passed(s *gherkin.Step, sd *godog.StepDef) {
	b.step(&stepResult{
		Result: "passed",
	})
	b.curElement.StepsPassed++
	b.curFeature.StepsPassed++
	b.StepsPassed++
}

// Skipped records a skipped step
func (b *CucumberFormatter) Skipped(s *gherkin.Step, sd *godog.StepDef) {
	b.step(&stepResult{
		Result: "skipped",
	})
	b.curStep.Result.Duration = nil
	b.curElement.StepsSkipped++
	b.curFeature.StepsSkipped++
	b.StepsSkipped++
}

// Undefined ...
func (b *CucumberFormatter) Undefined(s *gherkin.Step, sd *godog.StepDef) {
	b.step(&stepResult{
		Result: "undefined",
	})
	b.curStep.Result.Duration = nil
}

// Pending records steps that could not be matched to an implementation method
func (b *CucumberFormatter) Pending(s *gherkin.Step, sd *godog.StepDef) {
	b.step(&stepResult{
		Result: "pending",
	})
	b.curStep.Result.Duration = nil
	b.curStep.Match.Location = fmt.Sprintf("%s:%d", b.curStep.Match.Location, s.Location.Line)
	b.curElement.StepsPending++
	b.curFeature.StepsPending++
	b.StepsPending++
}

// Summary indicates the end of the test
func (b *CucumberFormatter) Summary() {
	var (
		jsonOutput []byte
		err        error
	)

	// Catch stats on the last scenario
	b.determineElementStatus()

	// Record total duration time
	b.RunTime = time.Now().Sub(b.startTime)

	if b.out != nil {
		if jsonOutput, err = json.Marshal(b.Features); err != nil {
			log.Printf(fmt.Sprintf(htmlErrMsg, "failed to marshal results to JSON: %s"), err)
			return
		}
		if _, err = b.out.Write(jsonOutput); err != nil {
			log.Printf(fmt.Sprintf(htmlErrMsg, "failed to write report file [%s]: %s"), common.ResultsJSONFile, err)
		}
	}
}
