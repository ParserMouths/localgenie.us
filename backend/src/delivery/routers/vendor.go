package routers

import "github.com/gofiber/fiber/v2"

func SetVendorRoutes(router *fiber.App) {
	router.Group("/user")
}
