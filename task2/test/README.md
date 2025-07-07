# Task2 BDD 測試與 HTML 報告

## 說明

此目錄包含 Task2 的 BDD 測試，使用 godog 執行測試，並透過專案共用的 HTML 報告產生器將測試結果轉換為 HTML 報告。

## 檔案結構

```
task2/test/
├── features/                  # BDD 特性檔案
│   └── chinese_chess.feature  # 中國象棋規則測試
├── main_test.go              # 測試主程式與 godog 設定
├── steps.go                  # 測試步驟實作
├── package.json              # Node.js 套件設定
└── reports/                  # 測試報告輸出目錄
    ├── cucumber.json         # JSON 格式測試結果
    └── cucumber-html-report.html # HTML 格式測試報告
```

**注意**: HTML 報告產生器現在使用專案共用的工具，位於 `../../tools/test-reporter/`

## 使用方法

### 1. 執行完整測試與報告產生

```bash
cd task2/test
npm run test
```

此命令會：
- 執行 godog BDD 測試
- 產生 JSON 格式的測試結果
- 自動轉換為 HTML 報告

### 2. 僅執行 BDD 測試

```bash
cd task2/test
npm run bdd-test
```

或直接使用 Go：

```bash
cd task2/test
go test -v
```

### 3. 僅產生 HTML 報告

如果已經有 `reports/cucumber.json` 檔案，可以只產生 HTML 報告：

```bash
cd task2/test
npm run html-report
```

或使用專案根目錄的 Makefile：

```bash
# 在專案根目錄執行
make html-report-task2
```

### 4. 檢視報告

HTML 報告產生後，可以在瀏覽器中開啟：

```bash
open reports/cucumber-html-report.html
```

## 報告內容

HTML 報告包含：
- 測試結果總覽
- 每個場景的詳細執行結果
- 通過/失敗的統計資訊
- 測試執行時間
- 測試環境資訊

## 使用 Makefile 執行

專案根目錄提供了 Makefile，可以更便於管理測試和報告產生：

```bash
# 在專案根目錄執行
make test-task2          # 只執行測試
make html-report-task2   # 只產生報告
make full-test-task2     # 執行測試並產生報告
make clean               # 清理所有報告檔案
```

## 技術細節

- **godog**: Golang 的 BDD 測試框架
- **共用報告產生器**: 使用 `tools/test-reporter/` 中的共用工具
- **cucumber-html-reporter**: 將 Cucumber JSON 格式轉換為 HTML 報告
- **輸出格式**: 同時產生 pretty 格式（終端機顯示）和 JSON 格式（報告使用）