# Memory 記憶命名規範

## 統一前綴系統

### 設計相關記憶 (DESIGN_)
- `design_api_{功能名稱}`: API 設計文件
- `design_database_{模組名稱}`: 資料庫設計文件
- `design_architecture_{系統名稱}`: 系統架構設計文件
- `design_workflow_{流程名稱}`: 工作流程設計文件
- `design_integration_{整合名稱}`: 系統整合設計文件

### 實作相關記憶 (IMPL_)
- `impl_progress_{功能名稱}`: 實作進度追蹤
- `impl_guide_{技術名稱}`: 實作指南文件
- `impl_pattern_{模式名稱}`: 實作模式文件

### 專案管理記憶 (PROJECT_)
- `project_structure`: 專案結構說明
- `project_workflow`: 開發工作流程
- `project_commands`: 常用開發指令
- `project_conventions`: 程式碼風格規範

### 技術文件記憶 (TECH_)
- `tech_framework_{框架名稱}`: 技術框架使用指南
- `tech_library_{函式庫名稱}`: 函式庫使用說明
- `tech_integration_{技術名稱}`: 技術整合說明

## 命名規則
1. **統一前綴**: 所有相關記憶使用統一前綴分類
2. **英文大寫**: 前綴和主要關鍵字使用英文大寫
3. **底線分隔**: 使用底線 `_` 分隔不同部分
4. **語意清晰**: 名稱應清楚表達記憶內容
5. **避免重複**: 新增前檢查是否已有相同類型記憶

## 更新流程
1. **檢查現有**: 使用 `list_memories` 檢查現有記憶
2. **識別類型**: 確定記憶所屬類別和前綴
3. **合併更新**: 如有相同類型記憶，進行合併更新
4. **避免重複**: 不建立重複或相似的記憶

## 清理原則
- 定期檢查過時記憶
- 合併相似內容記憶  
- 移除重複或無用記憶
- 保持記憶結構清晰
