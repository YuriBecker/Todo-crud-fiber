package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/yuribecker/todo-crud-fiber/config"
	"github.com/yuribecker/todo-crud-fiber/controller"
	"github.com/yuribecker/todo-crud-fiber/database"
	"github.com/yuribecker/todo-crud-fiber/middleware"
	"github.com/yuribecker/todo-crud-fiber/routes"
)

func init() {
	err := config.Load()
	if err != nil {
		log.Fatal("Error loading configs")
	}

	err = database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = database.Migrate()
	if err != nil {
		log.Fatal("Error migrating database")
	}
}

func main() {
	//setup fiber
	app := fiber.New()

	//setup middleware
	middleware.Init(app)

	//initialize app controller
	appController := controller.NewAppController(database.DB)

	//routes
	routes.SetupRoutes(app, appController)

	//start
	log.Fatal(app.Listen(":" + config.Get().API.Port))
}
