package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuribecker/todo-crud-fiber/controller"
)

func SetupRoutes(app *fiber.App, appController *controller.AppController) {
	mainRouter := app.Group("/api/v1")

	setupTodoRoutes(mainRouter, appController)
}
