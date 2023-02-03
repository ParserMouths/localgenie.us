package routers

import (
	"htf/src/delivery/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router *fiber.App, controller controllers.UserController) {
	r := router.Group("/user")
	r.Get("/test", controller.UserTest)
	r.Post("/signup", controller.UserSignUp)
	r.Get("/login", controller.UserSignIn)
}
