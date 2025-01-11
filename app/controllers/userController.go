package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"

	"github.com/gofiber/fiber/v2"
)

// Getuser func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} models.User
// @Router /v1/users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")
	if err := global.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "用户没找到",
		})
	}
	return c.JSON(fiber.Map{
		"message": "成功",
		"data":    user,
	})
}

// // Getusers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /v1/users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := global.Db.Find(&users).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "用户没找到",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   true,
		"message": "成功",
		"count":   len(users),
		"users":   users,
	})
}
