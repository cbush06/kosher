package suitecontext

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

// SuiteContext stores the status of test suite execution
// by means of registered callbacks with GoDog.
type SuiteContext struct {
	StepsPassed        int
	StepsFailed        int
	StepsSkipped       int
	StepsUndefined     int
	FeaturesPassed     int
	FeaturesFailed     int
	FeaturesIncomplete int
}

// CreateSuiteContext creates a new SuiteContext struct,
// registeres associated callbacks with GoDog, and returns
// a reference to the new struct.
func CreateSuiteContext(s *godog.Suite) *SuiteContext {
	s.BeforeFeature(beforeFeature)
	s.BeforeScenario(beforeScenario)
	s.BeforeStep(beforeStep)

	return &SuiteContext{}
}

func beforeFeature(f *gherkin.Feature) {}

func beforeScenario(s interface{}) {}

func beforeStep(s *gherkin.Step) {}

func afterFeature(f *gherkin.Feature) {}

func afterScenario(s interface{}, e error) {}

func afterStep(s *gherkin.Step, e error) {}
