package usecase_user

import (
	"context"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	config   *utils.Config
	db       *gorm.DB
	userRepo domain_user.Repository
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserUsecase(config *utils.Config, db *gorm.DB, userRepo domain_user.Repository) *userUsecase {
	return &userUsecase{
		config:   config,
		db:       db,
		userRepo: userRepo,
	}
}

func (handler *userUsecase) CreateUser(ctx context.Context, reqUser domain_user.User) (domain_user.User, error) {
	userID := uuid.New().String()
	if reqUser.IsVendor == "1" {

	}
	newUser := &domain_user.User{
		UserID:       userID,
		Username:     reqUser.Username,
		Firstname:    reqUser.Firstname,
		Lastname:     reqUser.Lastname,
		IsVendor:     reqUser.IsVendor,
		Latitude:     reqUser.Latitude,
		Longitude:    reqUser.Longitude,
		Email:        reqUser.Email,
		Password:     hashAndSalt(reqUser.Password),
		Subscription: reqUser.Subscription,
	}
	handler.userRepo.CreateUser(ctx, *newUser)
	return *newUser, nil
}

func (handler *userUsecase) VerifyUser(ctx context.Context, loginUser domain_user.LoginUser) (bool, string, domain_user.TokenReturn) {
	existingUser := handler.userRepo.GetUserFromUsername(ctx, loginUser.Username)
	if existingUser.UserID == "" {
		return false, "user not found", domain_user.TokenReturn{}
	}

	pwdVerified := comparePassword(existingUser.Password, loginUser.Password)
	if !pwdVerified {
		return false, "password not matched", domain_user.TokenReturn{}
	}

	tknReturn := &domain_user.TokenReturn{
		UserID: existingUser.UserID,
	}

	doStallIdExist := handler.userRepo.GetStallIdFromUserId(ctx, existingUser.UserID)
	if doStallIdExist != "" {
		tknReturn.StallID = doStallIdExist
	}

	return true, existingUser.UserID, *tknReturn
}

func (handler *userUsecase) GenerateAuthToken(ctx context.Context, loginUser domain_user.LoginUser) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claim := &Claims{
		Username: loginUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(handler.config.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func hashAndSalt(pwd string) string {
	bpwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bpwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal("Error hashing password")
	}
	return string(hash)
}

func comparePassword(hashed string, password string) bool {
	bHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bHash, []byte(password))
	if err != nil {
		return false
	}
	return true
}
