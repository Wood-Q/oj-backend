package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user_account body string true "UserAccount"
// @Param password body string true "Password"
// @Param user_role body string true "UserRole"
// @Success 200 {object} models.User
// @Router /api/v1/auth/sign/up [post]
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
	utils.SetupDatabase(c, models.User{})
	HashedPassword := utils.GeneratePassword(signUp.Password)
	//请求体解析为user结构体内容
	user := models.User{
		UserAccount:  signUp.UserAccount,
		UserPassword: HashedPassword,
		UserRole:     signUp.UserRole,
		UserName:     signUp.UserName,
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param user_account body string true "UserAccount"
// @Param password body string true "UserPassword"
// @Success 200 {string} status "ok"
// @Router /api/v1/auth/sign/in [post]
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
	utils.SetupDatabase(c, models.User{})
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
	//验证密码
	if !utils.ComparePasswords(foundedUser.UserPassword, signIn.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "密码错误",
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
	//将用户信息存入redis
	userInfo, err := json.Marshal(foundedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	err = global.RedisDb.Set(foundedUser.UserAccount, userInfo, time.Hour*72).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	//设置cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokens.Access,
		Expires:  time.Now().Add(72 * time.Hour),
		SameSite: fiber.CookieSameSiteStrictMode, // 强化跨站请求的安全性
	})

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

// GetLoginUser to parse jwt and get the login user
// @Description get current login user .
// @Summary get current login user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} any
// @Router /api/v1/auth/loginUser [get]
func GetLoginUser(c *fiber.Ctx) error {
	// 获取 token 中的 claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil || claims == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "token过期或没有token",
			"user":    fiber.Map{"userRole": "notLogin"},
		})
	}

	// 获取用户信息
	var user models.User
	if err := global.Db.Where("user_account=?", claims.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// 返回成功响应
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"过期时间为": time.Unix(int64(claims.Expires), 0).String(), // 格式化为正常时间
		"user":  user,
	})
}
