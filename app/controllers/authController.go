package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user_account body string true "UserAccount"
// @Param password body string true "Password"
// @Param user_role body string true "UserRole"
// @Success 200 {object} models.User
// @Router /v1/auth/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	//解析请求体
	signUp := &models.SignUp{}
	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	//数据库迁移
	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	//请求体解析为user结构体内容
	user := models.User{
		UserAccount:  signUp.UserAccount,
		UserPassword: signUp.Password,
		UserRole:     signUp.UserRole,
	}
	//创建用户
	if err := global.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	//返回200
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})

}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param user_account body string true "UserAccount"
// @Param password body string true "UserPassword"
// @Success 200 {string} status "ok"
// @Router /v1/auth/sign/in [post]
func UserSignIn(c *fiber.Ctx) error {
	//解析请求体
	signIn := &models.SignIn{}
	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	//数据库迁移
	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// 查找用户是否存在
	var foundedUser models.User // 定义一个变量用来存储查询到的用户
	if err := global.Db.Where("user_account=?", signIn.UserAccount).First(&foundedUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "用户名不存在",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "查询用户时出错",
		})
	}

	//生成token
	tokens, err := utils.GenerateNewTokens(foundedUser.UserAccount)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//返回200
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})

}
