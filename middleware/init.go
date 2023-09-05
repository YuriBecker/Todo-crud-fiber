package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	setupCORS(app)
	setupLogger(app)
	setupMonitor(app)
	configRecover(app)
}
