package controller_notification

import (
	"encoding/json"
	domain_notification "htf/src/internal/domain/notification"
	"htf/src/utils"

	"github.com/gofiber/fiber/v2"
)

type NotificationController interface {
	SendVapidKey(fiberHandler *fiber.Ctx) error
	CreateNewSubscription(fiberHandler *fiber.Ctx) error
	SendNotification(fiberHandler *fiber.Ctx) error
}

type controller struct {
	config       *utils.Config
	notification domain_notification.Usecase
}

func (c *controller) SendVapidKey(fiberHandler *fiber.Ctx) error {
	return fiberHandler.JSON(fiber.Map{
		"vapid_key": c.config.NotificationPublicKey,
	})
}

func (c *controller) CreateNewSubscription(fiberHandler *fiber.Ctx) error {
	var reqBody domain_notification.Subscription
	json.Unmarshal(fiberHandler.Body(), &reqBody)
	c.notification.CreateNewSubscription(fiberHandler.Context(), reqBody)
	return fiberHandler.JSON(fiber.Map{
		"test": "tdst",
	})
}

func (c *controller) SendNotification(fiberHandler *fiber.Ctx) error {
	var reqBody domain_notification.PushPayload
	json.Unmarshal(fiberHandler.Body(), &reqBody)

	err := c.notification.SendNotificationUsecase(fiberHandler.Context(), "")
	if err != nil {
		fiberHandler.JSON(fiber.Map{
			"message": "error",
		})
	}

	return fiberHandler.JSON(fiber.Map{
		"message": "success",
	})
}

func NewNotificationController(config *utils.Config, notificationUsecase domain_notification.Usecase) NotificationController {
	return &controller{
		config:       config,
		notification: notificationUsecase,
	}
}
