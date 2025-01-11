package utils

import (
	"OJ/pkg/configs" // 替换为您的实际模块路径
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// ParseJWT 解析和验证 JWT 访问令牌，并返回声明（claims）。

// 如果令牌无效或过期，返回相应的错误。
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// 从配置中获取 JWT 密钥
	secret := configs.AppConfig.JWT.Secret

	// 解析令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为预期的 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// if err != nil {
	// 	// 检查是否因为令牌过期导致的错误
	// 	var ve *jwt.ValidationError
	// 	if errors.As(err, &ve) {
	// 		if ve.Errors&jwt.ValidationErrorExpired != 0 {
	// 			return nil, fmt.Errorf("令牌已过期")
	// 		} else {
	// 			return nil, fmt.Errorf("令牌验证失败: %v", err)
	// 		}
	// 	}
	// 	return nil, fmt.Errorf("令牌解析错误: %v", err)
	// }

	// 断言令牌的声明类型为 jwt.MapClaims 并验证令牌有效性
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}
