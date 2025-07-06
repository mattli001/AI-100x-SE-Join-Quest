package test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/chinese_chess_zhTW.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	// 建立 Steps 實例
	steps := &Steps{}
	
	// 註冊步驟
	ctx.Step(`^棋盤為空，僅有一枚紅將於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅將於)
	ctx.Step(`^紅方將從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方將從移動至)
	ctx.Step(`^此步合法$`, steps.此步合法)
	ctx.Step(`^此步不合法$`, steps.此步不合法)
	ctx.Step(`^棋盤包含：$`, steps.棋盤包含)
	ctx.Step(`^棋盤為空，僅有一枚紅士於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅士於)
	ctx.Step(`^紅方士從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方士從移動至)
	ctx.Step(`^棋盤為空，僅有一枚紅車於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅車於)
	ctx.Step(`^紅方車從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方車從移動至)
	ctx.Step(`^棋盤為空，僅有一枚紅馬於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅馬於)
	ctx.Step(`^紅方馬從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方馬從移動至)
	ctx.Step(`^棋盤為空，僅有一枚紅炮於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅炮於)
	ctx.Step(`^紅方炮從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方炮從移動至)
	ctx.Step(`^棋盤為空，僅有一枚紅相於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅相於)
	ctx.Step(`^紅方相從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方相從移動至)
	ctx.Step(`^棋盤為空，僅有一枚紅兵於 \((\d+), (\d+)\)$`, steps.棋盤為空僅有一枚紅兵於)
	ctx.Step(`^紅方兵從 \((\d+), (\d+)\) 移動至 \((\d+), (\d+)\)$`, steps.紅方兵從移動至)
	ctx.Step(`^紅方立即獲勝$`, steps.紅方立即獲勝)
	ctx.Step(`^僅此吃子並不結束遊戲$`, steps.僅此吃子並不結束遊戲)
}

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features/chinese_chess_zhTW.feature"},
			Randomize: -1,
			Output:    colors.Colored(os.Stdout),
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}