package database

import (
	"context"
	domain_notification "htf/src/internal/domain/notification"
	domain_stall "htf/src/internal/domain/stall"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseClient(ctx context.Context, config *utils.Config) *gorm.DB {
	dsn := config.PostgresURI

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain_user.User{})
	db.AutoMigrate(&domain_notification.Subscription{})
	db.AutoMigrate(&domain_stall.Stall{})
	db.AutoMigrate(&domain_stall.Review{})
	return db
}
