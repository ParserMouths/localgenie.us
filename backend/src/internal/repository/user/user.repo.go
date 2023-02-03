package repo_user

import (
	"context"
	"fmt"
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
