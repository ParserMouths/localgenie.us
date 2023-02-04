package controller_stall

import (
	"encoding/json"
	"fmt"
	domain_stall "htf/src/internal/domain/stall"
	"htf/src/utils"

	// "time"

	"github.com/gofiber/fiber/v2"
)

type StallController interface {
	StallTest(fiberHandler *fiber.Ctx) (err error)
	CreateStall(fiberHandler *fiber.Ctx) (err error)
	UpdateStall(fiberHandler *fiber.Ctx) (err error)
	RemoveStall(fiberHandler *fiber.Ctx) (err error)
	QueryStall(fiberHandler *fiber.Ctx) (err error)
}

type controller struct {
	config *utils.Config
	stall  domain_stall.Usecase
}

func (c *controller) StallTest(fiberHandler *fiber.Ctx) (err error) {
	return fiberHandler.SendString("Hello from user")
}

func (c *controller) QueryStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) UpdateStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) RemoveStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) CreateStall(fiberHandler *fiber.Ctx) (err error) {
	var req_body domain_stall.Stall
	err = json.Unmarshal(fiberHandler.Body(), &req_body)
	if err != nil {
		fmt.Println(err)
	}
	newStall, err := c.stall.CreateStall(fiberHandler.Context(), req_body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(newStall)
	return fiberHandler.JSON(newStall)
}

func NewStallController(config *utils.Config, stallUseCase domain_stall.Usecase) StallController {
	return &controller{
		config: config,
		stall:  stallUseCase,
	}
}
