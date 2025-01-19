package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammad-rifqi/go_apps/database"
	"github.com/muhammad-rifqi/go_apps/database/migration"
	routers "github.com/muhammad-rifqi/go_apps/routes"
)

func main() {
	database.ConnectDB()
	migration.RunMigration()
	app := fiber.New()

	routers.RouterApp(app)

	app.Listen(":4000")
}
