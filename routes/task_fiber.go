package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kampus-kode-kode/go-railway/models"
)

func SetupTaskRoutes(app *fiber.App) {
	app.Get("/tasks", getTasks)
	app.Get("/tasks/:id", getTaskByID)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)
	app.Post("/tasks", createTask)
}

func createTask(ctx *fiber.Ctx) error {
	var newTask models.Task

	if err := ctx.BodyParser(&newTask); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	models.Tasks = append(models.Tasks, newTask)
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Task created"})
}

func getTasks(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"tasks": models.Tasks})
}

func getTaskByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	for _, task := range models.Tasks {
		if task.ID == id {
			return ctx.JSON(task)
		}
	}

	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}

func updateTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var updatedTask models.Task

	if err := ctx.BodyParser(&updatedTask); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for i, task := range models.Tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				models.Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				models.Tasks[i].Description = updatedTask.Description
			}
			return ctx.JSON(fiber.Map{"message": "Task updated"})
		}
	}

	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
}

func deleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	for i, val := range models.Tasks {
		if val.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return ctx.JSON(fiber.Map{"message": "Task removed"})
		}
	}

	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
}
