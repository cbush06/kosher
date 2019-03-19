package macros

import (
	"fmt"
	"os"
	"reflect"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/cbush06/kosher/fs"
	"github.com/spf13/afero"
)

// Macro wraps a macro's step text and substeps into a tidy package.
type Macro struct {
	Step     string
	Substeps *godog.Steps
}

// BuildMacros parses macro feature files and builds an array of
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
		for _, scenarioDef := range feature.ScenarioDefinitions {
			newMacro := &Macro{
				Step:     "",
				Substeps: &godog.Steps{},
			}

			substeps := []string{}

			switch scenario := scenarioDef.(type) {
			case *gherkin.Scenario:
				newMacro.Step = scenario.Name
				for _, substep := range scenario.Steps {
					stepText := substep.Text

					// If a DataTable or DocString is part
					// of the step, return an error
					if substep.Argument != nil {
						return nil, fmt.Errorf("Macros cannot contain data tables, docstrings, or scenario outlines")
					}
					substeps = append(substeps, stepText)
				}

				newMacro.Substeps = (*godog.Steps)(&substeps)
			default:
				return nil, fmt.Errorf("Only Scenarios permitted in macro Feature file. Encountered: %s", reflect.TypeOf(scenarioDef))
			}

			macros = append(macros, newMacro)
		}
	}

	return macros, nil
}
