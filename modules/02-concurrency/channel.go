package main

import "fmt"

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

// 計算 0 到 value 的總和
func sum(value int, ch chan int) {
	sum := 0
	for i := 0; i < value; i++ {
		sum += i
	}
	ch <- sum
}

func runChannelDemo() {
	fmt.Println("[channel] unbuffered 練習")
	// TODO: unbuffered：開一個 goroutine 送值，主程式收值並印出
	ch := make(chan int)
	go sum(10, ch)
	result := <-ch
	fmt.Println(result)

	fmt.Println("[channel] buffered 練習")
	// TODO: buffered：先送幾個值再收，觀察「沒人立刻來收」時是否還能送
	ch_buff := make(chan int, 2)
	ch_buff <- 1
	ch_buff <- 2

	// TODO: close(ch) 後用 for v := range ch 或 v, ok := <-ch 觀察行為
	close(ch_buff)           // 宣告：不會再送了
	for v := range ch_buff { // 把剩下的拿完就結束
		fmt.Println("v:", v)
	}
	fmt.Println("ch_buff is closed")
}
