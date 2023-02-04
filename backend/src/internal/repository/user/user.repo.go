package repo_user

import (
	"context"
	"fmt"
	domain_stall "htf/src/internal/domain/stall"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	config *utils.Config
	db     *gorm.DB
}

func NewUserRepository(config *utils.Config, db *gorm.DB) *UserRepository {
	return &UserRepository{
		config: config,
		db:     db,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user domain_user.User) error {
	fmt.Println(user)
	result := repo.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result.Error
}

func (repo *UserRepository) GetUserFromUsername(ctx context.Context, username string) domain_user.User {
	var user domain_user.User
	results := repo.db.Where("username = ?", username).First(&user)
	if results.Error != nil {
		return domain_user.User{}
	}
	return user
}

func (repo *UserRepository) GetStallIdFromUserId(ctx context.Context, userID string) string {
	var stall domain_stall.Stall
	results := repo.db.Table("stalls").Where("owner_id = ?", userID).First(&stall)
	if results.Error != nil {
		return ""
	}
	return stall.StallID
}
