package routes

// import (
// 	"net/http"

// 	"github.com/kampus-kode-kode/go-railway/models"
// 	"github.com/gin-gonic/gin"
// )

// func SetupTaskRoutes(router *gin.Engine) {
// 	router.GET("/tasks", getTasks)
// 	router.GET("/tasks/:id", getTaskByID)
// 	router.PUT("/tasks/:id", updateTask)
// 	router.DELETE("/tasks/:id", deleteTask)
// 	router.POST("/tasks", createTask)
// }

// func createTask(ctx *gin.Context) {
// 	var newTask models.Task

// 	if err := ctx.ShouldBindJSON(&newTask); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.Tasks = append(models.Tasks, newTask)
// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
// }

// func getTasks(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{"tasks": models.Tasks})
// }

// func getTaskByID(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	for _, task := range models.Tasks {
// 		if task.ID == id {
// 			ctx.JSON(http.StatusOK, task)
// 			return
// 		}
// 	}

// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// }

// func updateTask(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	var updatedTask models.Task

// 	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, task := range models.Tasks {
// 		if task.ID == id {
// 			if updatedTask.Title != "" {
// 				models.Tasks[i].Title = updatedTask.Title
// 			}
// 			if updatedTask.Description != "" {
// 				models.Tasks[i].Description = updatedTask.Description
// 			}
// 			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
// 			return
// 		}
// 	}

// 	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
// }

// func deleteTask(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	for i, val := range models.Tasks {
// 		if val.ID == id {
// 			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
// 			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
// 			return
// 		}
// 	}

// 	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
// }
