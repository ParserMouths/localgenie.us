package delivery

import (
	"context"
	"fmt"
	"htf/src/delivery/controllers"
	"htf/src/delivery/routers"
	"htf/src/internal/usecases"
	"htf/src/utils"

	"github.com/gofiber/fiber/v2"
)

func NewRestDelivery(ctx context.Context, config *utils.Config, useCases usecases.UseCases) {
	app := fiber.New()

	userController := controllers.NewUserController(config, useCases.User)

	routers.SetUserRoutes(app, userController)

	err := app.Listen(fmt.Sprintf(":%d", config.ServerPort))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on %d", config.ServerPort)
}
