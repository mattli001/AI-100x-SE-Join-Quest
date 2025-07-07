package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMain(m *testing.M) {
	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty,cucumber:reports/cucumber.json",
		Paths:  []string{"features"},
		Tags:   "", // 執行所有測試
	}

	status := godog.TestSuite{
		Name:                 "chinese_chess",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	// 測試套件初始化
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	// 中國象棋測試步驟
	steps := &ChineseChessSteps{}

	// Given steps
	ctx.Step(`^the board is empty except for a (.+) at \((\d+), (\d+)\)$`, steps.theBoardIsEmptyExceptForAPieceAt)
	ctx.Step(`^the board has:$`, steps.theBoardHas)

	// When steps
	ctx.Step(`^(.+) moves the (.+) from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, steps.playerMovesThePieceFromTo)

	// Then steps
	ctx.Step(`^the move is legal$`, steps.theMoveIsLegal)
	ctx.Step(`^the move is illegal$`, steps.theMoveIsIllegal)
	ctx.Step(`^(.+) wins immediately$`, steps.playerWinsImmediately)
	ctx.Step(`^the game is not over just from that capture$`, steps.theGameIsNotOverJustFromThatCapture)
}
