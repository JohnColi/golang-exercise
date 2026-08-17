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
func runRaceDemo() {
	fmt.Println("[Race] start")
	// TODO: unsafeCount()   — 不加鎖，印出錯誤（或不穩定）的結果
	// unsafeCount()
	// TODO: channelCount()  — 用 channel 收集增量，避免多 goroutine 同時寫同一變數
	channelCount()
	fmt.Println("[Race] end")
}

// 要保護的資料
type SafeCounter struct {
	mu sync.Mutex
	v  int
}

// #region 動物
// 鴨子
type Duck struct {
	Name string
}

// 雞
type Chicken struct {
	Name string
}

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
	// 每個 worker 把每次 +1 送進 channel；只有「接收端」改 total，避免多 goroutine 寫同一變數
	ch := make(chan int)
	var wg sync.WaitGroup
	washCount := 100
	feedCount := 100

	for i := 0; i < washCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1 // washing 完成一次
		}()
	}
	for i := 0; i < feedCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1 // feeding 完成一次
		}()
	}

	// 所有 worker 送完後關閉 channel，讓 range 能結束
	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for v := range ch {
		total += v // 單一 goroutine（main）負責加總
	}
	fmt.Println("[channel] final count:", total, "(expect", washCount+feedCount, ")")
}
