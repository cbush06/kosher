package macros

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/afero"

	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
	"github.com/cbush06/kosher/fs"
)

// Macro wraps a macro's step text and substeps into a tidy package.
type Macro struct {
	step     string
	substeps *godog.Steps
}

// BuildMacros parses macro feature fils and builds an array of
// Macros that can easily be added to the GoDog Suite as steps.
func BuildMacros(fs *fs.Fs) ([]*Macro, error) {
	var (
		features []*gherkin.Feature
		macros   []*Macro
		err      error
		errMsg   = "Error encountered while processing macros: %s"
	)

	if features, err = parseFeatures(fs); err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	if macros, err = createMacros(features); err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	return macros, nil
}

func parseFeatures(fs *fs.Fs) ([]*gherkin.Feature, error) {
	var (
		err          error
		errMsg       = "Error encountered discovering macro files: %s"
		featureFiles []string
		features     []*gherkin.Feature
	)

	if featureFiles, err = afero.Glob(fs.MacrosDir, "*.feature"); err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	for _, file := range featureFiles {
		var (
			featureFile afero.File
			newFeature  *gherkin.Feature
		)

		if featureFile, err = fs.MacrosDir.OpenFile(file, os.O_RDONLY, 0444); err != nil {
			return nil, fmt.Errorf(errMsg, err)
		}

		if newFeature, err = gherkin.ParseFeature(featureFile); err != nil {
			return nil, fmt.Errorf(errMsg, err)
		}

		features = append(features, newFeature)

		featureFile.Close()
	}

	return features, nil
}

func createMacros(features []*gherkin.Feature) ([]*Macro, error) {
	var macros []*Macro
	for _, feature := range features {
		newMacro := &Macro{
			step:     "",
			substeps: &godog.Steps{},
		}

		var substeps []string

		for _, step := range feature.ScenarioDefinitions {
			switch scenarioStep := step.(type) {
			case *gherkin.Scenario:
				newMacro.step = scenarioStep.Name
				for _, substep := range scenarioStep.Steps {
					stepText := substep.Text

					// If a DataTable or DocString is part of the step,
					// deserialize it and add it to the step's text
					if substep.Argument != nil {
						switch subStepArg := substep.Argument.(type) {
						case *gherkin.DataTable:
							for _, nextRow := range subStepArg.Rows {
								rowCells := []string{}
								for _, cell := range nextRow.Cells {
									rowCells = append(rowCells, cell.Value)
								}
								stepText += "\n" + strings.Join(rowCells, "|")
							}
						case *gherkin.DocString:
							stepText += "\n\"\"\"" + subStepArg.Content + "\n\"\"\""
						}
					}
					substeps = append(substeps, stepText)
				}

				newMacro.substeps = (*godog.Steps)(&substeps)
			default:
				return nil, fmt.Errorf("Only Scenarios permitted in macro Feature file. Encountered: %s", reflect.TypeOf(step))
			}
		}
		macros = append(macros, newMacro)
	}

	return macros, nil
}
