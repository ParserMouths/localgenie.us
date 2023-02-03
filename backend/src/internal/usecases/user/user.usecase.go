package usecase_user

import (
	"context"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userUsecase struct {
	config   *utils.Config
	db       *gorm.DB
	userRepo domain_user.Repository
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
	password := "test"
	newUser := &domain_user.User{
		UserID:    userID,
		Username:  reqUser.Username,
		Firstname: reqUser.Firstname,
		Lastname:  reqUser.Lastname,
		IsVendor:  reqUser.IsVendor,
		Location:  reqUser.Location,
		Email:     reqUser.Email,
		Password:  password,
	}
	handler.userRepo.CreateUser(ctx, *newUser)
	return *newUser, nil
}
