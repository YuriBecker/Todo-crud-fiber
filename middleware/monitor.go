package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func setupMonitor(app *fiber.App) {
	app.Get("/monitor", monitor.New())
}
