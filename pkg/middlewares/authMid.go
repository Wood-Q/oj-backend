package middlewares

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := utils.ExtractTokenMetadata(c)
		// log.Info("claims是", claims.UserID)
		if claims == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "没有token",
			})
		}
		if err != nil {
			log.Info("报错", err)
		}
		var user models.User
		if err := global.Db.Where("user_account=?", claims.UserID).First(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		if user.UserRole != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "权限不足",
			})
		}
		return c.Next()
	}
}

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := utils.ExtractTokenMetadata(c)
		if claims == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "没有token",
			})
		}
		if err != nil {
			log.Info("报错", err)
		}
		var user models.User
		if err := global.Db.Where("user_account=?", claims.UserID).First(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		currentUser:=user
		c.Locals("currentUser",&currentUser)
		return c.Next()
	}
}
