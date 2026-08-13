package main

// 練習 2.4：Select + timeout
//
// 目標：
// 1. 用 select 同時等待多個 channel
// 2. 實作簡單 timeout：若太久沒收到結果就走 timeout 分支
// 3. 想一想：多個 case 同時就緒時，select 怎麼選？default 有什麼用？
//
// 提示：
//   select {
//   case v := <-ch:
//       // 收到結果
//   case <-time.After(500 * time.Millisecond):
//       // timeout
//   }

func runSelectDemo() {
	// TODO: 建立一個 channel，在 goroutine 裡延遲一段時間後送出結果
	// TODO: 用 select 等待結果或 time.After 的 timeout
	// TODO: 調整延遲時間，分別觀察「準時收到」與「超時」兩種情況
}
