package test

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestDoubleElevenFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeDoubleElevenScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/double_eleven_promotion.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run double eleven feature tests")
	}
}

func InitializeDoubleElevenScenario(sc *godog.ScenarioContext) {
	ctx := &DoubleElevenTestContext{}
	ctx.RegisterDoubleElevenSteps(sc)
}
