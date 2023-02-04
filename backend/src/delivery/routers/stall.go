package routers

import (
	controller_stall "htf/src/delivery/controllers/stall"

	"github.com/gofiber/fiber/v2"
)

func SetStallRoutes(router *fiber.App, controller controller_stall.StallController) {
	r := router.Group("/stall")
	r.Get("/test", controller.StallTest)
	r.Post("/new", controller.CreateStall)
	r.Get("/:id", controller.QueryStall)
	r.Post("/review", controller.CreateReview)
	r.Post("/update/:id", controller.UpdateStall)
}
