package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammad-rifqi/go_apps/database"
	"github.com/muhammad-rifqi/go_apps/models/entity"
	"github.com/muhammad-rifqi/go_apps/models/req"
)

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.Database.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(req.UserReq)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Nama:  user.Nama,
		Email: user.Email,
	}

	if err := database.Database.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserControllerByID(c *fiber.Ctx) error {
	var User []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})
		return nil
	}

	if err := database.Database.Where("id = ?", id).First(&User).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Tidak DI Temukan",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    User,
	})
}
