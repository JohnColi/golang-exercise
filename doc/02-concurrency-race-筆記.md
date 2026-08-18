# Module 2.5：Race / Mutex / Channel — 問答筆記

對應練習：`modules/02-concurrency/race.go`

---

## 1. Deadlock：`all goroutines are asleep`

**現象：** `WaitGroup.Wait` 卡住，另兩個 goroutine 卡在 `chan send`。

**原因：**
- 無緩衝 channel 做了 `ch <- 1`，但程式裡**沒有任何人 receive**
- send 永遠等不到接收者 → worker 結束不了 → main 的 `wg.Wait()` 也永遠等

**解法（unsafe 示範）：** 先拿掉用不到的 channel；channel 留給後面的 `channelCount()`。

---

## 2. `count` 傳值 vs 傳指標

**問題：** 即使不死鎖，`final count` 仍可能一直是 `0`。

**原因：** `func f(count int)` 拿到的是**複本**，goroutine 裡的 `count++` 改不到 main 的變數。

**解法：** 傳 `*int`，用 `*count++`。  
（觀念同 WaitGroup：要傳 `&wg`，不能傳值。）

---

## 3. 印出指標內容

**現象：** `final count: 0x3b36bca20b0`（記憶體位址）

**原因：** 變數型別是 `*int`，直接 `Println(count)` 印的是位址。

**解法：** 印 `*count`。

---

## 4. 為什麼 `[Washing] final count: 10`？

**情境：** 開了很多個 washing goroutine，每個只做一次 `++` 就印「final count」。

**說明：** 那不是 washing「全部做完」的總和，只是**某一個** goroutine 印的當下，共享 `count` 剛好是 10。

真正最終結果：看 `wg.Wait()` 之後的那行 `final count`。

另外：多 goroutine 同時 `Println`，終端機行序 ≠ 實際發生順序。

---

## 5. 有 Mutex 了，為什麼同一個數字還印兩次？

**原因：** Mutex 只保護了臨界區；若寫成：

```go
mu.Lock()
*count++
mu.Unlock()
fmt.Println(*count) // 鎖外再讀
```

兩個 goroutine 可能在鎖外讀到**同一個快照**。

**Mutex 保證的是：** `++` 不交錯、最終總數正確。  
**不保證：** 鎖外印出來的數字彼此不重複。

**若要印「自己那次 +1 後」的值：**

```go
mu.Lock()
*count++
v := *count
mu.Unlock()
fmt.Println(v)
```

或把印出也放進鎖內。

---

## 6. `SafeCounter` 包住 Mutex

```go
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value // 在鎖內取出當下值再回傳
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
```

多 goroutine 共用同一個 `*SafeCounter`，透過 `Inc()` 加總。

---

## 7. Race Detector（`-race`）怎麼用？

**重點：** 不是程式裡呼叫的 API，而是執行時加旗標：

```bash
go run -race ./modules/02-concurrency
```

- 有鎖的 `SafeCounter`：通常**不會**報 DATA RACE  
- 無鎖同時寫同一個 `int`：會報 `WARNING: DATA RACE`

Windows 若出現 `CGO_ENABLED` / `gcc not found`，需安裝 MinGW/TDM-GCC 並啟用 CGO。  
沒有 `-race` 時，仍可觀察「最終 count 常小於期望值」。

---

## 8. `channelCount`：用 Channel 加總

**思路：** worker 不寫共享變數，只送增量；單一接收端負責加總。

```text
worker → ch <- 1
         …
wg.Wait() 後 close(ch)
main：for v := range ch { total += v }
```

| 作法 | 概念 |
|------|------|
| Mutex | 多人輪流鎖同一個變數 |
| Channel | 把增量交給一個人加總（share memory by communicating） |

期望結果（例如 100 wash + 100 feed）應穩定為 **200**。

---

## 對照表（練習目標）

| 函式 | 目的 | 典型結果 |
|------|------|----------|
| 無鎖同時 `count++` | 觀察 race / 少算 | 常 `< 期望值`；`-race` 會警告 |
| `SafeCounter` / Mutex | 保護共享變數 | 穩定正確 |
| Channel 收集再加總 | 避免多人寫同一變數 | 穩定正確 |
