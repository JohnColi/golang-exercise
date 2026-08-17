package main

import (
	"fmt"
	"sync"
)

// 練習 2.5：Race / Mutex / Channel 修正
//
// 目標：
// 1. 開多個 goroutine 同時對同一個計數器 +1，先不加鎖，觀察結果不正確
// 2. 用 sync.Mutex 修正
// 3. 再用 Channel（例如每個 goroutine 把增量送進 channel，由單一 goroutine 加總）修正
// 4. 比較兩種寫法的差異
//
// 建議用 -race 跑一次，看工具能不能抓到 race：
//
//	go run -race ./modules/02-concurrency
type SafeCounter struct {
	mu sync.Mutex
	v  int
}

// 鴨子
type Duck struct {
	Name string
}

// 雞
type Chicken struct {
	Name string
}

// #region 餵食和清洗
func goFeeding(name string, safeCount *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: 同上，但 +1 前後用 mutex.Lock() / Unlock()
	safeCount.mu.Lock()
	safeCount.v++
	defer safeCount.mu.Unlock()
	fmt.Println("[Feeding] ", name, ": final count:", safeCount.v)
}

func goWashing(name string, safeCount *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: 同上，但 +1 前後用 mutex.Lock() / Unlock()
	safeCount.mu.Lock()
	safeCount.v++
	defer safeCount.mu.Unlock()
	fmt.Println("[Washing] ", name, ": final count:", safeCount.v)
}

// #endregion 餵食和清洗

func runRaceDemo() {
	fmt.Println("[Race] start")
	// TODO: unsafeCount()   — 不加鎖，印出錯誤（或不穩定）的結果
	unsafeCount()

	// TODO: channelCount()  — 用 channel 收集增量，避免多 goroutine 同時寫同一變數
	channelCount()
	fmt.Println("[Race] end")
}

// 不加鎖，印出錯誤（或不穩定）的結果
func unsafeCount() {
	safeCount := SafeCounter{}
	// TODO: 啟動多個 goroutine，各自對同一個 int 做很多次 +1
	// TODO: 用 WaitGroup 等全部結束，印出最終 count（通常會小於預期）
	wg := sync.WaitGroup{}
	washCount := 100
	feedCount := 100

	for i := 0; i < washCount; i++ {
		wg.Add(1)
		go goWashing("chicken", &safeCount, &wg)
	}
	for i := 0; i < feedCount; i++ {
		wg.Add(1)
		go goFeeding("duck", &safeCount, &wg)
	}
	wg.Wait()

	fmt.Println("final count:", safeCount.v)
}

func channelCount() {
	// TODO: 每個 worker 把「完成次數」或每次 +1 送進 channel
	// TODO: 另開（或主程式）從 channel 收值做加總
}
