package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 練習 2.3：WaitGroup
//
// 目標：
// 1. 用 sync.WaitGroup 等待多個 goroutine 全部跑完
// 2. 搞清楚 Add / Done / Wait 的呼叫時機
// 3. 想一想：忘記 Done() 會發生什麼事？
//
// 提示：
//   var wg sync.WaitGroup
//   wg.Add(1)
//   go func() {
//       defer wg.Done()
//       // ...
//   }()
//   wg.Wait()

func goNumber1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goNumber_1")
}

func goNumber2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("goNumber_2 sleep 2 seconds")
}

func goNumber3(wg *sync.WaitGroup) {
	defer wg.Done()
	randomNumber := rand.Intn(4) + 1 // 1~10
	time.Sleep(time.Duration(randomNumber) * time.Second)
	fmt.Println("goNumber_3, sleep", randomNumber, "seconds")
}
func runWaitGroupDemo() {
	// TODO: 啟動 3 個 goroutine，各自印出自己的編號
	fmt.Println("[WaitGroup] start")
	var wg sync.WaitGroup
	wg.Add(3)
	// go worker(i, &wg) // ✅ 傳指標，大家共用同一個計數器
	// go worker(i, wg)  // ❌ 傳值，每個 goroutine 拿到的是「複製品」，改的不是同一個計數器
	go goNumber3(&wg)
	go goNumber2(&wg)
	go goNumber1(&wg)
	wg.Wait()
	fmt.Println("[WaitGroup] all done")
	// TODO: 用 WaitGroup 確保主程式等到全部結束才印 "all done"

}
