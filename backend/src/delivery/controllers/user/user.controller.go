package controller_user

import (
	"encoding/json"
	"fmt"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	UserTest(fiberHandler *fiber.Ctx) (err error)
	UserSignIn(fiberHandler *fiber.Ctx) (err error)
	UserSignUp(fiberHandler *fiber.Ctx) (err error)
}

type controller struct {
	config *utils.Config
	user   domain_user.UseCase
}

func (c *controller) UserTest(fiberHandler *fiber.Ctx) (err error) {
	return fiberHandler.SendString("Hello from user")
}

func (c *controller) UserSignIn(fiberHandler *fiber.Ctx) (err error) {
	var loginUser domain_user.LoginUser
	json.Unmarshal(fiberHandler.Body(), &loginUser)
	fmt.Println(loginUser)

	// check if password is same
	ok, prob, tknReturn := c.user.VerifyUser(fiberHandler.Context(), loginUser)
	if !ok {
		return fiberHandler.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": prob,
		})
	}

	tokenString, _ := c.user.GenerateAuthToken(fiberHandler.Context(), loginUser)

	tknReturn.Token = tokenString
	return fiberHandler.JSON(tknReturn)
}

func (c *controller) UserSignUp(fiberHandler *fiber.Ctx) (err error) {
	var req_body domain_user.User
	json.Unmarshal(fiberHandler.Body(), &req_body)
	newUser, err := c.user.CreateUser(fiberHandler.Context(), req_body)
	if err != nil {
		return err
	}
	return fiberHandler.JSON(newUser)
}

func NewUserController(config *utils.Config, userUseCase domain_user.UseCase) UserController {
	return &controller{
		config: config,
		user:   userUseCase,
	}
}
