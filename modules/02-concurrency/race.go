package main

// 練習 2.5：Race / Mutex / Channel 修正
//
// 目標：
// 1. 開多個 goroutine 同時對同一個計數器 +1，先不加鎖，觀察結果不正確
// 2. 用 sync.Mutex 修正
// 3. 再用 Channel（例如每個 goroutine 把增量送進 channel，由單一 goroutine 加總）修正
// 4. 比較兩種寫法的差異
//
// 建議用 -race 跑一次，看工具能不能抓到 race：
//   go run -race ./modules/02-concurrency

func runRaceDemo() {
	// TODO: unsafeCount()   — 不加鎖，印出錯誤（或不穩定）的結果
	// TODO: mutexCount()    — 用 Mutex 保護共享變數
	// TODO: channelCount()  — 用 channel 收集增量，避免多 goroutine 同時寫同一變數
}

func unsafeCount() {
	// TODO: 啟動多個 goroutine，各自對同一個 int 做很多次 +1
	// TODO: 用 WaitGroup 等全部結束，印出最終 count（通常會小於預期）
}

func mutexCount() {
	// TODO: 同上，但 +1 前後用 mutex.Lock() / Unlock()
}

func channelCount() {
	// TODO: 每個 worker 把「完成次數」或每次 +1 送進 channel
	// TODO: 另開（或主程式）從 channel 收值做加總
}
