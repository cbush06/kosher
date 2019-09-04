package websteps

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cbush06/kosher/report"
	"github.com/cbush06/kosher/resources/js"
	"github.com/cbush06/kosher/steps/steputils"
)

// Error maps directly to the JavaScript standard Error object
type Error struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

const axeScriptTemplate = `
    window.axeDone = false;
    window.axeError = null;
    window.axeResults = null;
    axe.run(
        document, 
        { 
            runOnly: { 
                type: "tag", 
                values: %s 
            } 
        },
        (err, results) => {
            window.axeError = err;
            window.axeResults = results;
            window.axeDone = true;
        }
    );
`

func iTestThePageForAccessibility(s *steputils.StepUtils) func() error {
	return func() error {
		var results []byte
		var err error

		pageTitle, _ := s.Page.Title()
		var axeResults = &report.AxeResults{
			Title: pageTitle,
		}

		// Load the AXE API
		s.Page.RunScript(string(js.AxeCore), nil, nil)

		// Run Axe
		if results, err = runAxeScanSync(s); err != nil {
			return err
		}

		// See if there  are any failures...most likely there are...
		if err = json.Unmarshal(results, &axeResults); err != nil {
			return fmt.Errorf("error encountered unmarshalling Axe results: %s", err)
		}

		// If at least one finding exceeds the threshold, fail this step
		// Also, count occurrences of each violation level
		var thresholdErr error
		impactThreshold := s.Settings.Settings.GetString("accessibility.impactThreshold")
		for _, violation := range axeResults.Violations {
			if thresholdErr == nil && report.AxeImpacts[violation.Impact] >= report.AxeImpacts[impactThreshold] {
				thresholdErr = fmt.Errorf("one or more accessibility findings exceeded the threshold [%s]", impactThreshold)
			}

			switch violation.Impact {
			case "minor":
				axeResults.MinorCount += len(violation.Nodes)
			case "moderate":
				axeResults.ModerateCount += len(violation.Nodes)
			case "serious":
				axeResults.SeriousCount += len(violation.Nodes)
			case "critical":
				axeResults.CriticalCount += len(violation.Nodes)
			}
		}

		for _, incomplete := range axeResults.Incomplete {
			axeResults.IncompleteCount += len(incomplete.Nodes)
		}

		// Store the result for inclusion in the final accessiblity report
		report.AddAxeResult(axeResults)

		return thresholdErr
	}
}

func runAxeScanSync(s *steputils.StepUtils) ([]byte, error) {
	var err error
	var axeError = new(Error)
	var rawData json.RawMessage

	// Get rule sets to run
	ruleSets := s.Settings.Settings.GetStringSlice("accessibility.ruleSets")

	// Write selected rule sets into script and run it in the browser
	ruleSetsJSON, _ := json.Marshal(ruleSets)
	axeScript := fmt.Sprintf(axeScriptTemplate, ruleSetsJSON)
	if err = s.Page.RunScript(axeScript, nil, nil); err != nil {
		return nil, fmt.Errorf("error encountered running the Axe tool: %s", err)
	}

	// Wait for Axe to finish
	axeScriptDone := false
	for !axeScriptDone {
		if err = s.Page.RunScript("return window.axeDone == true;", nil, &axeScriptDone); err != nil {
			return nil, fmt.Errorf("error encountered waiting for Axe tool to complete: %s", err)
		}
		time.Sleep(time.Duration(5) * time.Millisecond)
	}

	// Check for an Axe error
	if err = s.Page.RunScript("return window.axeError;", nil, &axeError); err != nil {
		return nil, fmt.Errorf("error encountered checking for Axe error: %s", err)
	}

	// If an error was reported, return that
	if axeError != nil {
		return nil, fmt.Errorf("error encountered by the Axe API: Name: %s; Message: %s", axeError.Name, axeError.Message)
	}

	// Get the results
	if err = s.Page.RunScript("return window.axeResults;", nil, &rawData); err != nil {
		return nil, fmt.Errorf("error encountered returning Axe scan results: %s", err)
	}

	return rawData, nil
}
