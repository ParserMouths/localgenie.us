package controller_user

import (
	"encoding/json"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("shlok-patel")

type UserController interface {
	UserTest(fiberHandler *fiber.Ctx) (err error)
	UserSignIn(fiberHandler *fiber.Ctx) (err error)
	UserSignUp(fiberHandler *fiber.Ctx) (err error)
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type controller struct {
	config *utils.Config
	user   domain_user.UseCase
}

func (c *controller) UserTest(fiberHandler *fiber.Ctx) (err error) {
	return fiberHandler.SendString("Hello from user")
}

func (c *controller) UserSignIn(fiberHandler *fiber.Ctx) (err error) {
	// get credentials
	// check if password is same
	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &Claims{
		Username: "Shlok",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fiberHandler.SendString("Cannot")
		return
	}
	return fiberHandler.SendString(tokenString)
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
