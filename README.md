# 微服務監控與熔斷系統

這是一個用於微服務架構的監控與管理系統，提供服務間請求追蹤、熔斷機制以及集中式日誌查詢功能。

## 功能特點

- **服務請求追蹤**：監控並記錄微服務之間的所有API呼叫與依賴關係
- **智能熔斷機制**：在服務異常時自動啟動熔斷，防止故障擴散
- **集中式日誌系統**：提供統一的介面查詢各服務的日誌資訊
- **即時監控面板**：視覺化展示服務健康狀態與性能指標

## 技術架構

本專案基於Go語言開發，採用現代化的微服務設計理念：

- 分散式追蹤：使用OpenTelemetry標準協議
- 熔斷實現：基於Circuit Breaker模式，支援自定義故障閾值
- 日誌收集：集成ELK或類似堆疊，支援結構化日誌
- 資料儲存：支援多種資料庫選項（關聯式/非關聯式）

## 快速開始

### 前置需求

- Go 1.19或更高版本
- Docker與Docker Compose（可選，用於本地部署）

### 安裝與運行

```bash
# 克隆專案
git clone https://github.com/yourusername/microservice-monitor.git
cd microservice-monitor

# 安裝依賴
go mod download

# 編譯專案
go build -o monitor ./cmd/app

# 運行服務
./monitor
```

## 配置與使用

### 基本配置

在`config.yaml`中設定您的服務參數：

```yaml
monitor:
  port: 8080
  logLevel: info
  
circuit_breaker:
  errorThreshold: 50      # 錯誤百分比閾值
  requestThreshold: 20    # 最小請求數閾值
  sleepWindow: 5000       # 熔斷後恢復嘗試時間(ms)
  
tracing:
  enabled: true
  samplingRate: 0.1       # 採樣率
```

### API文檔

啟動服務後，訪問 `http://localhost:8080/swagger/` 獲取完整API文檔。

## 開發與貢獻

歡迎提交Issue或Pull Request參與專案開發。

### 測試

```bash
go test ./...
```

## 授權

本專案採用MIT授權協議 - 詳見 [LICENSE](LICENSE) 文件。