# Module 3：Gin + RESTful API — 問答筆記

對應練習：`modules/03-gin/`

---

## 1. 什麼是 RESTful API？資源、HTTP method、狀態碼怎麼對應？

### HTTP 狀態碼與語意對應

狀態碼分類（第一位數字代表類別）：

| 類別 | 意義 |
|------|------|
| 2xx | 成功 |
| 3xx | 重新導向 |
| 4xx | 客戶端錯誤（請求本身有問題） |
| 5xx | 伺服器錯誤 |

常見對應：

| 狀態碼 | 意義 | 常見情境 |
|--------|------|----------|
| 200 OK | 成功，回傳資料 | GET、PUT、PATCH 成功 |
| 201 Created | 成功建立資源 | POST 成功，通常會在 header 帶 Location 指向新資源 URL |
| 204 No Content | 成功但無回傳內容 | DELETE 成功 |
| 400 Bad Request | 請求格式錯誤 | 缺欄位、型別錯誤 |
| 401 Unauthorized | 未認證 | 沒帶 token 或 token 無效 |
| 403 Forbidden | 已認證但無權限 | 權限不足 |
| 404 Not Found | 資源不存在 | `/users/999`（該 id 不存在） |
| 405 Method Not Allowed | 該資源不支援此 method | 對唯讀資源用 DELETE |
| 409 Conflict | 資源狀態衝突 | 重複建立、版本衝突 |
| 422 Unprocessable Entity | 格式對但邏輯驗證失敗 | 欄位值不合法 |
| 500 Internal Server Error | 伺服器內部錯誤 | 未預期的例外 |

---

## 2. Gin 解決了 `net/http` 的哪些不方便之處？

（待填）

---

## 3. 怎麼在 Gin 裡定義路由、綁定 JSON、回傳 JSON？

（待填：路由定義、c.JSON 回傳）

### 常見誤區（綁定 JSON / 驗證）

#### 1. Body 只能讀一次

`ShouldBindJSON` 會讀取並「消耗」request body（底層是 stream）。若需要讀兩次（例如先 log 原始 body，再 bind），要用會快取 body 的 `c.ShouldBindBodyWith()`，否則第二次讀會是空的。

```go
// 錯誤示範：第二次 bind 會失敗，因為 body 已經被讀完
c.ShouldBindJSON(&user1)
c.ShouldBindJSON(&user2) // 出錯或拿到空值
```

#### 2. `binding:"required"` 對「零值」的判斷要注意型別

例如 `bool` 加 `required` 時，`false` 會被當成零值而驗證失敗，常常不是預期行為。要小心設計欄位型別，或改用指標（`*bool`）搭配驗證邏輯。

#### 3. 只驗證 body，不驗證 URL 參數或 query string

`ShouldBindJSON` 只處理 body：

| 要綁什麼 | 用哪個 |
|----------|--------|
| JSON body | `ShouldBindJSON` |
| 路徑參數（如 `:id`） | `ShouldBindUri` |
| query string | `ShouldBindQuery` |

`ShouldBind()` 會依 `Content-Type` 自動選綁定方式，但明確指定（如 `ShouldBindJSON`）通常比較清楚、不易出錯。

---

## 4. 路由分組（route group）的用途是什麼？

（待填）
