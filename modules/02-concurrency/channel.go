package main

// 練習 2.2：Channel
//
// 目標：
// 1. 建立 unbuffered / buffered channel，各送收一次值
// 2. 觀察什麼時候會 block（卡住）
// 3. 關閉 channel，並用 ok 判斷是否還有值可收
//
// 提示：
//   ch := make(chan int)      // unbuffered
//   ch := make(chan int, 2)   // buffered，容量 2
//   v, ok := <-ch             // ok == false 表示 channel 已關閉且空了

func runChannelDemo() {
	// TODO: unbuffered：開一個 goroutine 送值，主程式收值並印出
	// TODO: buffered：先送幾個值再收，觀察「沒人立刻來收」時是否還能送
	// TODO: close(ch) 後用 for v := range ch 或 v, ok := <-ch 觀察行為
}
