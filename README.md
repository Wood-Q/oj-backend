# Fiber后端模板

## 简介
基于Golang的Fiber框架的后端模板，包含了一些常用的中间件和功能模块，方便快速搭建后端服务。

## 所需go包
```
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/swagger
go get github.com/golang-jwt/jwt/v5
go get github.com/redis/go-redis
go get golang.org/x/crypto
go get github.com/swaggo/swag
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/spf13/viper
go get github.com/stretchr/testify
```

## 所需go工具
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

   - [github.com/swaggo/swag](https://github.com/swaggo/swag) 自动生成api文档
   - [github.com/securego/gosec](https://github.com/securego/gosec) 检查go代码的安全性问题
   - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) 检查go代码是否符合最佳实践
   - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) go代码静态分析工具

## 整体介绍

### ./app

**Folder with business logic only**. 包含业务逻辑和与应用程序核心逻辑相关的模块，通常是控制器、模型

- `./app/controllers` 是用于控制器controllers
- `./app/models` 用于定义项目gorm所需的模型

### ./docs

**Folder with API Documentation**. 用于存储swagger生成的api文档

### ./pkg

**Folder with project specific functionality**. 存放项目的通用功能模块，适用于业务无关的代码逻辑。

- `./pkg/configs` 项目的配置文件或配置相关逻辑，方便集中管理和修改。
- `./pkg/middleware` 存放 Fiber 框架的中间件，比如认证、日志、错误处理等。
- `./pkg/routes` 定义了所有 API 的路由配置，将请求与具体控制器绑定。
- `./pkg/enums` 集中定义和管理常量
- `./pkg/utils` 存储工具函数或通用方法，例如日期处理、加密解密等逻辑。

### ./platform

**Folder with platform-level logic**. 与平台相关的逻辑，可能包含缓存、数据库迁移和其他底层支持代码。

- `./platform/cache` 缓存相关的配置和逻辑代码。
- `./platform/database` 存放数据库连接逻辑或初始化代码。


