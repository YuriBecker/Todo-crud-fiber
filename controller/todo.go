package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yuribecker/todo-crud-fiber/model"
	"gorm.io/gorm"
)

type todoController struct {
	DB *gorm.DB
}

func newTodoController(db *gorm.DB) *todoController {
	return &todoController{
		DB: db,
	}
}

func (t *todoController) FindAll(c *fiber.Ctx) error {
	var todos []model.Todo

	t.DB.Find(&todos)

	return c.JSON(todos)
}

func (t *todoController) FindById(c *fiber.Ctx) error {
	var todo model.Todo

	t.DB.First(&todo, c.Params("id"))

	return c.JSON(todo)
}

func (t *todoController) Insert(c *fiber.Ctx) error {
	todo := new(model.Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	t.DB.Create(&todo)

	return c.Status(http.StatusCreated).JSON(todo)
}

func (t *todoController) Update(c *fiber.Ctx) error {
	var todo model.Todo

	t.DB.First(&todo, c.Params("id"))

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	t.DB.Save(&todo)

	return c.Status(http.StatusOK).JSON(todo)
}

func (t *todoController) Delete(c *fiber.Ctx) error {
	var todo model.Todo

	t.DB.Delete(&todo, c.Params("id"))

	c.Status(http.StatusNoContent)

	return nil
}
