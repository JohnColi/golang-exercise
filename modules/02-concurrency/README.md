# Module 2：並發（Concurrency）

對應 `golang學習項目.md` 的 Module 2。

## 檔案說明

| 檔案 | 對應 | 用途 |
|------|------|------|
| `main.go` | 入口 | 依序串起各練習 |
| `goroutine.go` | 2.1 | 啟動 goroutine、主程式結束會砍掉未完成工作 |
| `channel.go` | 2.2 | buffered / unbuffered、關閉與接收 |
| `waitgroup.go` | 2.3 | `Add` / `Done` / `Wait` |
| `select.go` | 2.4 | `select` + timeout |
| `race.go` | 2.5 + 綜合練習 | 不加鎖 → Mutex → Channel 三種加總 |

## 怎麼跑

```powershell
go run ./modules/02-concurrency
```

檢查 race condition：

```powershell
go run -race ./modules/02-concurrency
```

## 建議順序

1. `goroutine.go` → 在 `main.go` 取消註解 `runGoroutineDemo()`
2. `channel.go` → `runChannelDemo()`
3. `waitgroup.go` → `runWaitGroupDemo()`
4. `select.go` → `runSelectDemo()`
5. `race.go` → `runRaceDemo()`（大綱的綜合練習）

一個觀念跑通再進下一個；做完後回大綱把對應問題勾起來。
