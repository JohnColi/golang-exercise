package main

import "fmt"

// 練習：用記憶體存資料（先不接資料庫）
//
// 目標：
// 1. 用 slice 或 map 保存資源
// 2. 實作基本 CRUD：Create / List / Get / Update / Delete
// 3. 之後如果要加並發安全，可再考慮 Mutex（Module 2）

// TODO: 定義 in-memory store（例如 var todos []Todo 或 map[int]Todo）
// TODO: 實作 CRUD 函式（純資料層，不要直接碰 Gin context）

var todos []Todo

// CRUD 操作的函式
func CreateTodo(todo Todo) {
	todos = append(todos, todo)
}

// #region Get
func GetTodos() []Todo {
	return todos
}

func GetTodoByID(id int) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	// 如果沒找到，回傳空的 Todo
	fmt.Println("Todo not found")
	return Todo{}
}

func GetTodoByTitle(title string) Todo {
	for _, todo := range todos {
		if todo.Title == title {
			return todo
		}
	}
	return Todo{}
}

func GetTodoByCompleted(completed bool) []Todo {
	var result []Todo
	for _, todo := range todos {
		if todo.Completed == completed {
			result = append(result, todo)
		}
	}
	return result
}

// #endregion

func DeleteTodoByID(id int) bool {
	isFound := false
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			isFound = true
			return isFound
		}
	}

	return isFound
}

func UpdateTodoByID(id int, todo Todo) (Todo, bool) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = todo.Title
			todos[i].Completed = todo.Completed
			return todos[i], true
		}
	}
	fmt.Println("Todo not found")
	return Todo{}, false
}
