package routes

import (
	"OJ/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Post("/auth/sign/up",controllers.UserSignUp)
	route.Post("/auth/sign/in",controllers.UserSignIn)

}