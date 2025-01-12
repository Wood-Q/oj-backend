package routes

import (
	"OJ/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicAuthRoutes(a *fiber.App) {

	route := a.Group("/api/v1/auth")

	route.Post("/sign/up", controllers.UserSignUp)
	route.Post("/sign/in", controllers.UserSignIn)
	route.Get("/loginUser", controllers.GetLoginUser)

}

func PublicUserRoutes(a *fiber.App) {
	route := a.Group("/api/v1/users")
	route.Get("/", controllers.GetUsers)
	route.Get("/:id", controllers.GetUserById)
}

func PublicQuestionRoutes(a *fiber.App) {
	route := a.Group("/api/v1/questions")
	route.Get("/", controllers.GetQuestions)
	route.Get("/:question_id", controllers.GetQuestion)
	route.Delete("/:question_id", controllers.DeleteQuestion)
	route.Post("/", controllers.CreateQuestion)
	route.Put("/:question_id", controllers.UpdateQuestion)
}

func PublicQuestionSubmitRoutes(a *fiber.App) {
	route := a.Group("/api/v1/questions")
	route.Get("/", controllers.GetQuestionSubmits)
	route.Get("/:question_id", controllers.GetQuestionSubmit)
	route.Delete("/:question_id", controllers.DeleteQuestionSubmit)
	route.Post("/", controllers.CreateQuestionSubmit)
	route.Put("/:question_id", controllers.UpdateQuestionSubmit)
}
