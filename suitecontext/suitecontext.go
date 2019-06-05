package suitecontext

import (
	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
)

const (
	passed = iota
	failed
	undefined
)

// SuiteContext stores the status of test suite execution
// by means of registered callbacks with GoDog.
type SuiteContext struct {
	StepsCount               int
	StepsPassed              int
	StepsFailed              int
	StepsSkipped             int
	StepsUndefined           int
	ScenariosCount           int
	ScenariosPassed          int
	ScenariosFailed          int
	ScenariosUndefined       int
	FeaturesCount            int
	FeaturesPassed           int
	FeaturesFailed           int
	FeaturesUndefined        int
	currentFeature           *gherkin.Feature
	currentFeatureStatus     int
	currentScenarioStepCount int
}

// CreateSuiteContext creates a new SuiteContext struct,
// registers associated callbacks with GoDog, and returns
// a reference to the new struct.
func CreateSuiteContext(s *godog.Suite) *SuiteContext {
	suite := &SuiteContext{}

	s.BeforeFeature(suite.beforeFeature)
	s.BeforeScenario(suite.beforeScenario)
	s.BeforeStep(suite.beforeStep)
	s.AfterFeature(suite.afterFeature)
	s.AfterScenario(suite.afterScenario)
	s.AfterStep(suite.afterStep)

	return suite
}

func (suite *SuiteContext) beforeFeature(f *gherkin.Feature) {
	suite.currentFeature = f
	suite.FeaturesCount++
}

func (suite *SuiteContext) beforeScenario(s interface{}) {
	suite.ScenariosCount++
	suite.currentScenarioStepCount = 0
}

func (suite *SuiteContext) beforeStep(s *gherkin.Step) {
	suite.StepsCount++
}

func (suite *SuiteContext) afterFeature(f *gherkin.Feature) {
	switch suite.currentFeatureStatus {
	case passed:
		suite.FeaturesPassed++
	case failed:
		suite.FeaturesFailed++
	case undefined:
	default:
		suite.FeaturesUndefined++
	}
}

func (suite *SuiteContext) afterScenario(s interface{}, e error) {
	switch e {
	case godog.ErrUndefined:
		suite.ScenariosUndefined++
		suite.StepsUndefined++
		suite.StepsSkipped += suite.getTotalStepsOfScenario(s) - suite.currentScenarioStepCount
		suite.currentFeatureStatus = undefined
	case nil:
		suite.ScenariosPassed++
		suite.currentFeatureStatus = passed
	default:
		suite.ScenariosFailed++
		suite.currentFeatureStatus = failed
	}
}

func (suite *SuiteContext) getTotalStepsOfScenario(s interface{}) int {
	totalSteps := 0

	switch s.(type) {
	case *gherkin.Scenario:
		totalSteps = len((s.(*gherkin.Scenario)).Steps)
	case *gherkin.ScenarioOutline:
		totalSteps = len((s.(*gherkin.ScenarioOutline)).Steps)
	}

	if suite.currentFeature.Background != nil {
		totalSteps += len(suite.currentFeature.Background.Steps)
	}

	return totalSteps
}

func (suite *SuiteContext) afterStep(s *gherkin.Step, e error) {
	switch e {
	case nil:
		suite.StepsPassed++
		suite.currentScenarioStepCount++
	default:
		suite.StepsFailed++
	}
}
