package routes

import (
	"OJ/app/controllers"
	"OJ/pkg/middlewares"

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
	route.Use(middlewares.CheckAuth())
	route.Get("/", controllers.GetQuestions)
	route.Get("/:question_id", controllers.GetQuestionByID)
	route.Get("/dividePage/questions", controllers.GetQuestionsByPage)
	route.Post("/", controllers.CreateQuestion)
	route.Use(middlewares.CheckAdmin())
	route.Delete("/:question_id", controllers.DeleteQuestion)
	route.Put("/:question_id", controllers.UpdateQuestion)
}

func PublicQuestionSubmitRoutes(a *fiber.App) {
	route := a.Group("/api/v1/questionsSubmit")
	route.Use(middlewares.CheckAuth())
	route.Get("/", controllers.GetQuestionSubmits)
	// route.Get("/:question_id", controllers.GetQuestionSubmit)
	route.Post("/", controllers.CreateQuestionSubmit)
	route.Delete("/:question_id", controllers.DeleteQuestionSubmit)
	// route.Put("/:question_id", controllers.UpdateQuestionSubmit)
}
