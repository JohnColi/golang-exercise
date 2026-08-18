# Module 3：Web 框架與 RESTful API（Gin）

對應 `golang學習項目.md` 的 Module 3。

## 檔案說明

| 檔案 | 用途 |
|------|------|
| `main.go` | 程式入口，啟動 Gin server |
| `models.go` | 資源（resource）的資料結構，例如 Todo |
| `store.go` | 記憶體內資料存取（slice / map，先不接資料庫） |
| `handlers.go` | CRUD handler：綁定 JSON、回傳 JSON |
| `routes.go` | 定義路由與路由分組（route group） |

## 怎麼跑

先安裝依賴（專案根目錄）：

```powershell
go get github.com/gin-gonic/gin
```

再啟動：

```powershell
go run ./modules/03-gin
```

預設監聽 `http://localhost:8080`。

## 建議順序

1. 先讀大綱問題：REST 資源、HTTP method、狀態碼怎麼對應
2. 完成 `models.go` + `store.go`（決定要 CRUD 什麼資源）
3. 完成 `handlers.go`（綁定 JSON、回傳 JSON、對應狀態碼）
4. 完成 `routes.go`（路由 + route group）
5. 在 `main.go` 串起來，用 curl / Postman / 瀏覽器驗證 CRUD
6. 回大綱把對應問題勾起來

筆記可寫在 `doc/03-gin-筆記.md`。
