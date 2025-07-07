# BDD 測試報告產生器

這是一個共用的 BDD 測試結果 HTML 報告產生工具。

## 功能特色

- 將 Cucumber JSON 測試結果轉換為 HTML 報告
- 支援可設定的參數
- 統一的報告樣式和格式
- 支援多專案使用

## 使用方法

### 基本使用

```bash
npm install
node generate-html-report.js
```

### 使用參數

```bash
node generate-html-report.js --jsonFile reports/cucumber.json --output reports/report.html --projectName "我的專案" --feature "功能名稱"
```

### 參數說明

- `--jsonFile`: JSON 測試結果檔案路徑（預設：reports/cucumber.json）
- `--output`: HTML 報告輸出路徑（預設：reports/cucumber-html-report.html）
- `--projectName`: 專案名稱（預設：BDD 測試專案）
- `--version`: 版本號（預設：1.0.0）
- `--environment`: 測試環境（預設：Development）
- `--feature`: 功能描述（選填）

## 在其他專案中使用

1. 在您的專案中新增 npm script：
```json
{
  "scripts": {
    "html-report": "node ../../tools/test-reporter/generate-html-report.js --projectName '您的專案名稱'"
  }
}
```

1. 或直接在 Makefile 中使用：
  ```makefile
  html-report:
    cd tools/test-reporter && npm install
    node tools/test-reporter/generate-html-report.js --projectName "您的專案名稱"
  ```
