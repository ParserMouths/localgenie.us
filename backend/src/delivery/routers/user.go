package routers

import (
	controller_user "htf/src/delivery/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router *fiber.App, controller controller_user.UserController) {
	r := router.Group("/user")
	r.Get("/test", controller.UserTest)
	r.Post("/signup", controller.UserSignUp)
	r.Post("/login", controller.UserSignIn)
}

func SetRestrictedUserRoutes(router *fiber.App, controller controller_user.UserController) {
	r := router.Group("/user")
	r.Get("/test/restricted", controller.UserTest)
}
