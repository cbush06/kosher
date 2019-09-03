package report

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/resources/reporttemplates"
	"github.com/spf13/afero"
)

const axeErrMsg = "error encountered while generated Axe report: %s"

// AxeImpacts lists the violation impact categories in order of severity
var AxeImpacts = map[string]int{
	"minor":    1,
	"moderate": 2,
	"serious":  3,
	"critical": 4,
}

// AxeReport encapsulates details of all Axe scans for a Kosher session
type AxeReport struct {
	ProjectName     string        `json:"projectName"`
	OS              string        `json:"os"`
	AppVersion      string        `json:"appVersion"`
	Timestamp       string        `json:"timestamp"`
	Environment     string        `json:"environment"`
	Browser         string        `json:"browser"`
	RuleSets        []string      `json:"ruleSets"`
	ImpactThreshold string        `json:"impactThreshold"`
	AxeVersion      string        `json:"axeVersion"`
	AxeScans        []*AxeResults `json:"scans"`
	ViolationsCount int           `json:"violationsCount"`
	IncompleteCount int           `json:"incompleteCount"`
	MinorCount      int           `json:"minorCount"`
	ModerateCount   int           `json:"moderateCount"`
	SeriousCount    int           `json:"seriousCount"`
	CriticalCount   int           `json:"criticalCount"`
}

// AxeResults is the top-level results object returned by a call to axe.run(...)
type AxeResults struct {
	Title           string              `json:"title"`
	TestEngine      *AxeTestEngine      `json:"testEngine"`
	TestEnvironment *AxeTestEnvironment `json:"testEnvironment"`
	ToolOptions     *AxeToolOptions     `json:"toolOptions"`
	URL             string              `json:"url"`
	Timestamp       string              `json:"timestamp"`
	Violations      []AxeRuleFindings   `json:"violations"`
	Incomplete      []AxeRuleFindings   `json:"incomplete"`
	MinorCount      int                 `json:"minorCount"`
	ModerateCount   int                 `json:"moderateCount"`
	SeriousCount    int                 `json:"seriousCount"`
	CriticalCount   int                 `json:"criticalCount"`
	IncompleteCount int                 `json:"incompleteCount"`
}

// AxeTestEngine stores metadata about the version of the axe-core API used
type AxeTestEngine struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// AxeTestEnvironment stores metadata about the browser and screen size used during testing
type AxeTestEnvironment struct {
	OrientationAngle int    `json:"orientationAngle"`
	OrientationType  string `json:"orientationType"`
	UserAgent        string `json:"userAgent"`
	WindowHeight     int    `json:"windowHeight"`
	WindowWidth      int    `json:"windowWidth"`
}

// AxeToolOptions stores options and settings passed to Axe for the test run
type AxeToolOptions struct {
	RunOnly *AxeRunOnly `json:"runOnly"`
}

// AxeRunOnly stores limits set on the rule sets applied for the test
type AxeRunOnly struct {
	Type   string   `json:"runOnly"`
	Values []string `json:"values"`
}

// AxeRuleFindings stores the rule details and findings for each rule executed against the page.
type AxeRuleFindings struct {
	Description string     `json:"description"`
	Help        string     `json:"help"`
	HelpURL     string     `json:"helpUrl"`
	ID          string     `json:"id"`
	Impact      string     `json:"impact"`
	Tags        []string   `json:"tags"`
	Nodes       []AxeNodes `json:"nodes"`
}

// AxeNodes stores the findings of a rule and relates these findings to DOM nodes on a page.
type AxeNodes struct {
	HTML           string         `json:"html"`
	Impact         string         `json:"impact"`
	Target         []interface{}  `json:"target"`
	FailureSummary string         `json:"failureSummary"`
	Any            []AxeNodeCheck `json:"any"`
	All            []AxeNodeCheck `json:"all"`
	None           []AxeNodeCheck `json:"none"`
}

// GetPrettyFailureSummary returns a "prettified" version of the failure summary
func (a *AxeNodes) GetPrettyFailureSummary() template.HTML {
	fs := "<strong>" + a.FailureSummary
	fs = strings.ReplaceAll(fs, ":\n  ", ":</strong><ul><li>")
	fs = strings.ReplaceAll(fs, "\n  ", "</li><li>")
	fs = strings.ReplaceAll(fs, "\n\n", "</li></ul><strong>")
	fs = fs + "</li></ul>"
	return template.HTML(fs)
}

// AxeNodeCheck stores finding details for a specific DOM node.
type AxeNodeCheck struct {
	ID           string                `json:"id"`
	Impact       string                `json:"impact"`
	Message      string                `json:"message"`
	Data         interface{}           `json:"data"`
	RelatedNodes []AxeNodeRelatedNodes `json:"relatedNodes"`
}

// AxeNodeRelatedNodes relates a node's finding to other nodes involved in that finding (e.g. nodes that share a duplicate ID).
type AxeNodeRelatedNodes struct {
	Target []string `json:"target"`
	HTML   string   `json:"html"`
}

// Axe is the report structure that contains the summary and results of all Axe scans from the current session
var Axe *AxeReport

// NewAxeReport creates and stores a new Axe report. This should only be called once per application run, as it creates a global Axe report.
func NewAxeReport(s *config.Settings) {
	Axe = &AxeReport{
		ProjectName:     s.Settings.GetString("projectName"),
		OS:              fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH),
		AppVersion:      s.Settings.GetString("appVersion"),
		Timestamp:       time.Now().Format(time.RFC850),
		Environment:     s.Settings.GetString("environment"),
		Browser:         s.Settings.GetString("driver"),
		RuleSets:        s.Settings.GetStringSlice("accessibility.ruleSets"),
		ImpactThreshold: s.Settings.GetString("accessibility.impactThreshold"),
	}
}

// AddAxeResult stores an AxeResults entry for output in the final report
func AddAxeResult(result *AxeResults) {
	Axe.AxeScans = append(Axe.AxeScans, result)
	Axe.AxeVersion = result.TestEngine.Version

	Axe.ViolationsCount += (result.MinorCount + result.ModerateCount + result.SeriousCount + result.CriticalCount)
	Axe.MinorCount += result.MinorCount
	Axe.ModerateCount += result.ModerateCount
	Axe.SeriousCount += result.SeriousCount
	Axe.CriticalCount += result.CriticalCount
	Axe.IncompleteCount += result.IncompleteCount
}

// Process generates the Axe report
func (a *AxeReport) Process(s *config.Settings) error {
	var (
		templ      *template.Template
		fileHandle afero.File
		err        error
		writeLen   int
		jsonOutput []byte
	)

	const (
		resultsHTMLFile = common.AxeResultsHTMLFile
		resultsJSONFile = common.AxeResultsJSONFile
	)

	funcMap := template.FuncMap{
		"last": func(s []interface{}) interface{} {
			return s[len(s)-1]
		},
	}

	// Do not generate a report if there are no items to report
	if a.ViolationsCount < 1 && a.IncompleteCount < 1 {
		return nil
	}

	templ, _ = template.New("AxeBootstrap").Funcs(funcMap).Parse(reporttemplates.GetAxeTemplate())

	// write HTML report file
	filePath, _ := s.FileSystem.ResultsDir.RealPath(resultsHTMLFile)
	if fileHandle, err = s.FileSystem.ResultsDir.OpenFile(resultsHTMLFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to open Axe results file [%s]: %s"), filePath, err)
	}
	if err = templ.Execute(fileHandle, Axe); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to generate Axe report file [%s]: %s"), filePath, err)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to close Axe report file [%s]: %s"), filePath, err)
	}

	// write JSON results file
	filePath, _ = s.FileSystem.ResultsDir.RealPath(resultsJSONFile)
	if fileHandle, err = s.FileSystem.ResultsDir.OpenFile(resultsJSONFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to open Axe results file [%s]: %s"), filePath, err)
	}
	if jsonOutput, err = json.Marshal(Axe); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to marshal to JSON: %s"), err)
	}
	if writeLen, err = fileHandle.Write(jsonOutput); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to write report file [%s]: %s"), filePath, err)
	} else if writeLen < len(jsonOutput) {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to write all bytes of report file [%s]"), filePath)
	}
	if err = fileHandle.Close(); err != nil {
		return fmt.Errorf(fmt.Sprintf(axeErrMsg, "failed to close report file [%s]: %s"), filePath, err)
	}

	return nil
}
