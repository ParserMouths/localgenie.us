package repo_stall

import (
	"context"
	"fmt"
	domain_stall "htf/src/internal/domain/stall"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"

	"gorm.io/gorm"
)

type StallRepository struct {
	config *utils.Config
	db     *gorm.DB
}

func NewStallRepository(config *utils.Config, db *gorm.DB) *StallRepository {
	return &StallRepository{
		config: config,
		db:     db,
	}
}

func (repo *StallRepository) CreateStall(ctx context.Context, stall domain_stall.Stall) error {
	result := repo.db.Create(&stall)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return result.Error
}

func (repo *StallRepository) UpdateStall(ctx context.Context, stallId string, stall domain_stall.Stall) error {
	result := repo.db.Where("stall_id = ?", stallId).Updates(stall)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return result.Error
}

func (repo *StallRepository) RemoveStall(ctx context.Context, stall domain_stall.Stall) error {
	fmt.Println(stall)
	result := repo.db.Delete(&stall, stall.StallID)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return result.Error
}

func (repo *StallRepository) GetStallFromStallID(ctx context.Context, stallID string) (domain_stall.Stall, error) {
	var stall domain_stall.Stall
	result := repo.db.Table("stalls").Where("stall_id = ?", stallID).First(&stall)
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain_stall.Stall{}, result.Error
	}
	return stall, nil
}

func (repo *StallRepository) GetUsersAroundStall(ctx context.Context, latitude float32, longitude float32) ([]domain_user.User, error) {
	var users []domain_user.User
	result := repo.db.Table("users").Where("latitude BETWEEN ? AND ?", latitude-0.01, latitude+0.01).Where("longitude BETWEEN ? AND ?", longitude-0.01, longitude+0.01).Find(&users)
	if result.Error != nil {
		return []domain_user.User{}, result.Error
	}
	return users, nil
}
