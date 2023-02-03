package routers

import "github.com/gofiber/fiber/v2"

func SetStallRoutes(router *fiber.App) {
	router.Group("/stall")
}
