package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammad-rifqi/go_apps/controllers"
)

func RouterApp(c *fiber.App) {
	c.Get("/", controllers.UserControllerShow)
}
