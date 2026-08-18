package main

// 練習：路由與路由分組（route group）
//
// 目標：
// 1. 把 HTTP method + path 對到 handler
// 2. 用 Group 把同一前綴的路由收在一起（例如 /api/v1）
// 3. 想一想：route group 之後對 middleware（Module 4）有什麼幫助？

// TODO: 實作 setupRouter()，回傳 *gin.Engine
// 範例方向：
//   r := gin.Default()
//   api := r.Group("/api/v1")
//   {
//       api.GET("/todos", listTodos)
//       api.POST("/todos", createTodo)
//       api.GET("/todos/:id", getTodo)
//       api.PUT("/todos/:id", updateTodo)
//       api.DELETE("/todos/:id", deleteTodo)
//   }
//   return r
