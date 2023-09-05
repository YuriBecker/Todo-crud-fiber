package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuribecker/todo-crud-fiber/controller"
)

func setupTodoRoutes(r fiber.Router, ac *controller.AppController) {
	app := r.Group("/todo")

	app.Get("/:id", ac.Todo.FindById)
	app.Get("", ac.Todo.FindAll)
	app.Post("", ac.Todo.Insert)
	app.Put("/:id", ac.Todo.Update)
	app.Delete("/:id", ac.Todo.Delete)
}
