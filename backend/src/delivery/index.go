package delivery

import (
	"context"
	"fmt"
	controller_notification "htf/src/delivery/controllers/notification"
	controller_user "htf/src/delivery/controllers/user"
	"htf/src/delivery/routers"
	"htf/src/internal/usecases"
	"htf/src/utils"

	"github.com/gofiber/fiber/v2"
)

func NewRestDelivery(ctx context.Context, config *utils.Config, useCases usecases.UseCases) {
	app := fiber.New()

	userController := controller_user.NewUserController(config, useCases.User)
	notificationController := controller_notification.NewNotificationController(config, useCases.Notification)

	routers.SetUserRoutes(app, userController)
	routers.SetNotificationRoutes(app, notificationController)

	err := app.Listen(fmt.Sprintf(":%d", config.ServerPort))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on %d", config.ServerPort)
}
