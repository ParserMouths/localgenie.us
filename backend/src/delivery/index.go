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
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewRestDelivery(ctx context.Context, config *utils.Config, useCases usecases.UseCases) {
	app := fiber.New()
	app.Use(cors.New())

	userController := controller_user.NewUserController(config, useCases.User)
	notificationController := controller_notification.NewNotificationController(config, useCases.Notification)

	routers.SetUserRoutes(app, userController)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("shlok-patel"),
	}))
	routers.SetRestrictedUserRoutes(app, userController)
	routers.SetNotificationRoutes(app, notificationController)

	err := app.Listen(fmt.Sprintf(":%v", config.ServerPort))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on %v", config.ServerPort)
}
