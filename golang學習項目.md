# Golang 學習專案：從並發到 RESTful API

> 一個適合分享給 Go 初學者的自學路線圖。每個階段只列出「該搞懂的問題」與「該做的練習」，**先不寫答案**，鼓勵自己動手查文件、寫程式碼驗證。

---

## 專案目標

透過一個「小型 Web 服務」當作主線專案，把 Go 的核心觀念（物件導向風格、並發、Web 框架）串起來練習，而不是零散地看語法。

**建議的主線專案**：用 Gin 寫一個簡單的 RESTful API（例如待辦事項 / 短網址服務 / 留言板），並在其中加入並發處理（例如批次任務、背景工作）。

---

## Module 1：Go 的物件導向風格

Go 沒有 `class`，但依然可以做到封裝、組合、多型。

**要搞懂的問題：**
- [ ] Go 用什麼取代 `class`？
- [ ] 沒有繼承（inheritance）的情況下，Go 怎麼做到「共用行為」或「擴充功能」？（提示：關鍵字是「組合」而不是「繼承」）
- [ ] `struct` 內嵌（embedding）跟傳統物件導向的繼承，本質上有什麼不同？
- [ ] interface 在 Go 裡扮演什麼角色？跟其他語言的 interface 有何不同（例如不需要顯式宣告 implements）？

**練習建議：**
- 寫兩個 struct，讓其中一個「內嵌」另一個，觀察方法是怎麼被繼承下來的。
- 設計一個 interface，讓不同的 struct 都實作它，體會 Go 的「鴨子型別」。

---

## Module 2：並發（Concurrency）— 這次筆記的核心

Go 的招牌特色。建議照下面順序理解，一個觀念沒懂，後面會卡關。

### 2.1 Goroutine
- [ ] Goroutine 是什麼？跟作業系統的 thread 有什麼差別？
ans:
thread 是 OS 的單位；goroutine 是 Go 自己排的「更輕的併發單位」。

- [ ] 怎麼啟動一個 goroutine？它跟主程式的執行順序關係是什麼？
- [ ] 為什麼主程式結束時，還沒跑完的 goroutine 會直接被砍掉？

### 2.2 Channel
- [ ] Channel 是拿來做什麼的？跟共享變數的方式有什麼不同的思維（"share memory by communicating"）？
- [ ] buffered channel 跟 unbuffered channel 差在哪？什麼時候會 block？

ans:
**Buffered vs unbuffered，什麼時候會 block？**
| | Unbuffered `make(chan T)` | Buffered `make(chan T, n)` |
|---|---|---|
| 容量 | 0 | n |
| 送（send） | 一定要有人立刻收，否則 block | buffer 未滿可直接送；滿了才 block |
| 收（receive） | 一定要有人立刻送，否則 block | buffer 有資料可直接收；空了才 block |

- [ ] 怎麼判斷一個 channel 已經被關閉（closed）？

ans:
``` golang
v, ok := <-ch
if !ok {
    // channel 已關閉，而且裡面沒東西了
}
for v := range ch {
    // 處理 v
}
// 走到這裡表示 ch 已關閉且資料收完
```

### 2.3 WaitGroup
- [ ] 為什麼需要 WaitGroup？它解決了什麼問題？
- [ ] `Add` / `Done` / `Wait` 分別在什麼時機呼叫？
- [ ] 如果忘記呼叫 `Done()` 會發生什麼事？

### 2.4 Select
- [ ] `select` 的用途是什麼？跟 `switch` 有何不同？
- [ ] 如果同時有多個 channel 都準備好了，`select` 怎麼決定要走哪一個？
- [ ] `select` 裡的 `default` 分支有什麼作用？

### 2.5 Mutex
- [ ] 什麼情況下光靠 channel 不夠，需要用到 Mutex？
- [ ] `sync.Mutex` 跟 `sync.RWMutex` 的差別是什麼？
- [ ] Race condition 是什麼？Go 有沒有工具可以幫忙抓出來？（提示：`-race`）

**練習建議：**
- 寫一個小程式，開多個 goroutine 同時對同一個變數做加總，先不加鎖，觀察錯誤結果；再分別用 Mutex 和 Channel 兩種方式修正，比較兩種寫法的差異。
- 用 `select` + channel 實作一個簡單的 timeout 機制。

---

## Module 3：Web 框架與 RESTful API（Gin）

- [ ] 什麼是 RESTful API？資源（resource）、HTTP method、狀態碼之間怎麼對應？
- [ ] Gin 這個框架解決了 Go 內建 `net/http` 的哪些不方便之處？
- [ ] 怎麼在 Gin 裡定義路由（route）、綁定 JSON 參數、回傳 JSON？
- [ ] 路由分組（route group）的用途是什麼？

**練習建議：**
- 用 Gin 寫出一組基本 CRUD API（新增／查詢／更新／刪除），先不套用資料庫，用記憶體內的 slice 或 map 存資料即可。

---

## Module 4：Middleware

- [ ] Middleware 是什麼？在請求（request）處理流程中，它站在哪個位置？
- [ ] 常見的 middleware 用途有哪些？（例如：記錄 log、驗證身份、處理 CORS）
- [ ] 在 Gin 裡，middleware 是怎麼被串接（chain）起來的？`c.Next()` 的作用是什麼？

**練習建議：**
- 自己寫一個簡單的 logging middleware，印出每個請求的方法、路徑、處理時間。
- 試著寫一個驗證用的 middleware（例如檢查 header 是否帶了某個 token），並套用到部分路由上。

---

## 建議學習順序

```
Module 1（物件導向風格）
      ↓
Module 2（並發：Goroutine → Channel → WaitGroup → Select → Mutex）
      ↓
Module 3（Gin + RESTful API）
      ↓
Module 4（Middleware）
      ↓
整合：把 Module 2 的並發技巧，用在 Module 3 的 API 裡
      （例如：某個 API 需要同時打好幾個外部請求，用 goroutine + channel 收集結果）
```

---

## 給分享對象的提醒

這份大綱刻意只列「問題」與「練習方向」，沒有附上程式碼答案。
建議搭配：
- Go 官方文件（[go.dev/doc](https://go.dev)）
- Go by Example（範例式教學）
- 實際動手把每個問題寫成小程式驗證，比看答案更有效。