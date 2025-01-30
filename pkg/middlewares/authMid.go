package middlewares

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// claims, err := utils.ExtractTokenMetadata(c)
		// // log.Info("claims是", claims.UserID)
		// if claims == nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error":   true,
		// 		"message": "没有token",
		// 	})
		// }
		// if err != nil {
		// 	log.Info("报错", err)
		// }
		// var user models.User
		// if err := global.Db.Where("user_account=?", claims.UserID).First(&user).Error; err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error":   true,
		// 		"message": err.Error(),
		// 	})
		// }
		currentUser, ok := c.Locals("currentUser").(*models.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "用户信息无效",
			})
		}
		if currentUser.UserRole != "admin" {
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
		// var user models.User
		// if err := global.Db.Where("user_account=?", claims.UserID).First(&user).Error; err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error":   true,
		// 		"message": err.Error(),
		// 	})
		// }
		userInfo, err := global.RedisDb.Get(claims.UserID).Result()
		if err == redis.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token 过期或无效",
			})
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		// log.Info("userInfo是", userInfo)
		c.Locals("currentUser", &userInfo)
		return c.Next()
	}
}
