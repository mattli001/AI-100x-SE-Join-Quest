# BDD 測試和報告產生 Makefile

.PHONY: help install-reporter test-task2 test-task2-2 test-all html-report-task2 html-report-task2-2 html-report-all clean

help:
	@echo "可用的指令："
	@echo "  install-reporter  - 安裝共用的測試報告產生器"
	@echo "  test-task2        - 執行 task2 的 BDD 測試"
	@echo "  test-task2-2      - 執行 task2-2 的 BDD 測試"
	@echo "  test-all          - 執行所有 task 的 BDD 測試"
	@echo "  html-report-task2 - 產生 task2 的 HTML 報告"
	@echo "  html-report-task2-2 - 產生 task2-2 的 HTML 報告"
	@echo "  html-report-all   - 產生所有 task 的 HTML 報告"
	@echo "  clean            - 清理所有產生的報告檔案"

install-reporter:
	@echo "安裝共用測試報告產生器..."
	cd tools/test-reporter && npm install

test-task2:
	@echo "執行 task2 BDD 測試..."
	cd task2/test && go test -v

test-task2-2:
	@echo "執行 task2-2 BDD 測試..."
	cd task2-2/test && go test -v

test-all: test-task2 test-task2-2

html-report-task2: install-reporter
	@echo "產生 task2 HTML 報告..."
	cd task2/test && node ../../tools/test-reporter/generate-html-report.js \
		--projectName "Task2 中國象棋" \
		--feature "中國象棋 BDD 測試" \
		--version "1.0.0"

html-report-task2-2: install-reporter
	@echo "產生 task2-2 HTML 報告..."
	cd task2-2/test && node ../../tools/test-reporter/generate-html-report.js \
		--projectName "Task2-2 中國象棋進階" \
		--feature "中國象棋進階 BDD 測試" \
		--version "1.0.0"

html-report-all: html-report-task2 html-report-task2-2

full-test-task2: test-task2 html-report-task2
	@echo "Task2 完整測試和報告產生完成"

full-test-task2-2: test-task2-2 html-report-task2-2
	@echo "Task2-2 完整測試和報告產生完成"

full-test-all: test-all html-report-all
	@echo "所有測試和報告產生完成"

clean:
	@echo "清理報告檔案..."
	rm -rf task2/test/reports/*.html
	rm -rf task2-2/test/reports/*.html
	rm -rf task2/test/reports/*.json
	rm -rf task2-2/test/reports/*.json
	@echo "清理完成"

show-reports:
	@echo "可用的報告："
	@find task2/test/reports task2-2/test/reports -name "*.html" 2>/dev/null || echo "尚未產生任何報告"