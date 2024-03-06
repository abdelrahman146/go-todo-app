package main

import (
	"todo-app/controllers"
	"todo-app/providers"

	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/todos", controllers.FetchAllTodos)
	r.GET("/todos/:id", controllers.FetchSingleTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.PUT("/todos/:id/mark-done", controllers.MarkAsDone)
	r.PUT("/todos/:id/mark-udone", controllers.MarkAsUndone)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	return r
}

func main() {
	providers.InitDatabase()
	r := setup()
	r.Run(":8080")
}
