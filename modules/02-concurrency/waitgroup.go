package main

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

func runWaitGroupDemo() {
	// TODO: 啟動 3 個 goroutine，各自印出自己的編號
	// TODO: 用 WaitGroup 確保主程式等到全部結束才印 "all done"
}
