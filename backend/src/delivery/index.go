package delivery

import (
	"context"
	"fmt"
	controller_notification "htf/src/delivery/controllers/notification"
	controller_stall "htf/src/delivery/controllers/stall"
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
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Authorization,Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	userController := controller_user.NewUserController(config, useCases.User)
	notificationController := controller_notification.NewNotificationController(config, useCases.Notification)
	stallController := controller_stall.NewStallController(config, useCases.Stall)

	routers.SetUserRoutes(app, userController)
	routers.SetNotificationRoutes(app, notificationController)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JwtSecret),
	}))
	routers.SetRestrictedUserRoutes(app, userController)
	routers.SetStallRoutes(app, stallController)

	err := app.Listen(fmt.Sprintf(":%v", config.ServerPort))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on %v", config.ServerPort)
}
