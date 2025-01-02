.PHONY: clean critic security lint test build run swag migrate.up migrate.down migrate.force docker.run docker.stop docker.network docker.fiber.build docker.fiber docker.postgres docker.redis docker.stop.fiber docker.stop.postgres docker.stop.redis


# 静态代码检查
critic:
	gocritic check -enableAll ./...


# 安全检查
security:
	gosec ./...


# 代码格式和质量检查
lint:
	golangci-lint run ./...

# 生成 Swagger 文档
swag:
	swag init --parseDependency --parseInternal --output ./docs/swagger
