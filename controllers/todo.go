package controllers

import (
	"todo-app/models"
	"todo-app/providers"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(400, gin.H{"success": false, "error": err.Error()})
		return
	}
	db := providers.GetDB()
	db.Create(&newTodo)
	c.JSON(200, gin.H{"success": true, "data": newTodo})
}

func FetchAllTodos(c *gin.Context) {
	var todos []models.Todo
	db := providers.GetDB()
	db.Find(&todos)
	c.JSON(200, gin.H{"success": true, "data": todos})
}

func FetchSingleTodo(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")
	db := providers.GetDB()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(404, gin.H{"success": false, "error": "No todo found"})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": todo})
}

func UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	todoID := c.Param("id")
	db := providers.GetDB()
	db.First(&updatedTodo, todoID)
	if updatedTodo.ID == 0 {
		c.JSON(404, gin.H{"success": false, "error": "No todo found"})
		return
	}
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(400, gin.H{"success": false, "error": err.Error()})
		return
	}
	db.Save(&updatedTodo)
	c.JSON(200, gin.H{"success": true, "data": updatedTodo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")
	db := providers.GetDB()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(404, gin.H{"success": false, "error": "No todo found"})
		return
	}
	db.Delete(&todo)
	c.JSON(200, gin.H{"success": true, "data": todo})
}

func MarkAsDone(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")
	db := providers.GetDB()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(404, gin.H{"success": false, "error": "No todo found"})
		return
	}
	todo.Status = true
	db.Save(&todo)
	c.JSON(200, gin.H{"success": true, "data": todo})
}

func MarkAsUndone(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")
	db := providers.GetDB()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(404, gin.H{"success": false, "error": "No todo found"})
		return
	}
	todo.Status = true
	db.Save(&todo)
	c.JSON(200, gin.H{"success": true, "data": todo})
}

func DeleteAllTodos(c *gin.Context) {
	db := providers.GetDB()
	db.Exec("DELETE FROM todos")
	c.JSON(200, gin.H{"success": true, "data": "All todos deleted"})
}
