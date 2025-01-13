package main

import (
	_ "OJ/docs" // load API Docs files (Swagger)
	"OJ/pkg/configs"
	"OJ/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs.InitConfig()
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	routes.PublicUserRoutes(app)
	routes.PublicQuestionSubmitRoutes(app)
	routes.PublicQuestionRoutes(app)
	routes.PublicAuthRoutes(app)
	routes.SwaggerRoute(app)
	app.Listen(configs.AppConfig.App.Port)
}
