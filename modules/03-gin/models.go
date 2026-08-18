package main

// 練習：定義 API 要操作的「資源」（resource）
//
// 目標：
// 1. 想清楚這個資源對應哪個 URL（例如 /todos、/todos/:id）
// 2. JSON 欄位用 struct tag 標好（`json:"..."`）
// 3. 對照大綱：資源、HTTP method、狀態碼怎麼對應

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// 暫時註解
// type Note struct {
// 	ID      int    `json:"id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// 暫時註解
// type ShortURL struct {
// 	ID        int       `json:"id"`
// 	URL       string    `json:"url"`
// 	ShortURL  string    `json:"short_url"`
// 	CreatedAt time.Time `json:"created_at"`
// }
