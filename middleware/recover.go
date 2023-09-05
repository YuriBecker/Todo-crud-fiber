package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func configRecover(app *fiber.App) {
	app.Use(recover.New())
}
