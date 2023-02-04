package routers

import (
	controller_notification "htf/src/delivery/controllers/notification"

	"github.com/gofiber/fiber/v2"
)

func SetNotificationRoutes(router *fiber.App, controller controller_notification.NotificationController) {
	r := router.Group("/notification")
	r.Get("/key", controller.SendVapidKey)
	r.Post("/subscription/create", controller.CreateNewSubscription)
	r.Post("/send", controller.SendNotification)
}
