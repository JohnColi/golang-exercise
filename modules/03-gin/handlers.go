package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 練習：Gin handler — 綁定 JSON、回傳 JSON
//
// 目標：
// 1. 用 c.ShouldBindJSON / c.BindJSON 接 request body
// 2. 用 c.JSON 回傳資料與狀態碼
// 3. 對照 REST：POST 建、GET 查、PUT/PATCH 改、DELETE 刪

// TODO: 實作各 CRUD handler，例如：
func createTodo(context *gin.Context) {
	var todo Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CreateTodo(todo)
	// 回傳 created 狀態碼和 todo 資料
	// HTTP/1.1 201 Created Content-Type: application/json
	// {"id":1,"title":"...","completed":false}
	context.JSON(http.StatusCreated, todo)
}

func listTodos(context *gin.Context) {
	todos := GetTodos()
	context.JSON(http.StatusOK, todos)
}

func getTodoByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := GetTodoByID(id)
	context.JSON(http.StatusOK, todo)
}

func deleteTodoByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DeleteTodoByID(id)
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func updateTodo(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo Todo
	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, ok := UpdateTodoByID(id, todo)
	if !ok {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, updated)
}
