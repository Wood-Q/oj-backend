package utils

import (
	"OJ/pkg/configs"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID  string
	Expires float64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := VerifyToken(c)

	if err != nil {
		log.Info("报错", err)
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID := claims["id"].(string)

		var expires float64
		switch v := claims["exp"].(type) {
		case float64:
			expires = v
		case string:
			expiresVal, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Info("无法转换 exp 字段:", err)
				return nil, err
			}
			expires = expiresVal
		default:
			log.Info("exp 字段类型无效")
			return nil, err
		}

		return &TokenMetadata{
			UserID:  userID,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func ExtractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	cleanedToken := strings.Trim(bearToken, "[]")

	return cleanedToken
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(configs.AppConfig.JWT.Secret), nil
}
