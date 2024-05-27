package main

import (
	"github.com/gofiber/fiber/v2"
	routers "github.com/muhammad-rifqi/go_apps/routes"
)

func main() {
	app := fiber.New()

	routers.RouterApp(app)

	app.Listen(":4000")
}
