package main

// 練習 2.1：Goroutine
//
// 目標：
// 1. 用 go 啟動一個（或多個）goroutine
// 2. 觀察主程式與 goroutine 的執行順序（可能交錯、不一定誰先）
// 3. 想一想：為什麼主程式一結束，還沒跑完的 goroutine 會被砍掉？
//
// 提示：可搭配 time.Sleep 暫時讓主程式多等一下（之後會用 WaitGroup 取代這種寫法）

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func runGoroutineDemo() {
	// TODO: 印出 "main start"
	fmt.Println("=== runGoroutineDemo start ===")

	// TODO: 用 go 啟動一個函式，在裡面印出 "hello from goroutine"
	go say("hello world")

	time.Sleep(1 * time.Second) // 等 goroutine 跑完（練習用）
	// TODO: 印出 "main end"
	fmt.Println("=== runGoroutineDemo end ===")
	// TODO: 觀察：若主程式太快結束，goroutine 的訊息可能看不到
}
