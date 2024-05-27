package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammad-rifqi/go_apps/database"
	"github.com/muhammad-rifqi/go_apps/entity"
)

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.Database.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}
