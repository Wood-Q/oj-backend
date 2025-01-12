package utils

import (
	"OJ/pkg/global"

	"github.com/gofiber/fiber/v2"
)

func SetupDatabase(c *fiber.Ctx, model interface{}) error {
	// 自动迁移数据库，确保表结构与模型一致
	if err := global.Db.AutoMigrate(model); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return nil
}
