package domain_stall

import (
	"context"
	domain_user "htf/src/internal/domain/user"
)

type Stall struct {
	StallID     string `json:"stall_id" gorm:"primaryKey"`
	OwnerID     string `json:"owner_id" gorm:"references:UserID"`
	StallName   string `json:"stall_name"`
	IsOpen      int    `json:"is_open"`
	CreatedAt   string `json:"created_at"`
	Rating      string `json:"rating"`
	LastActive  string `json:"last_active"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offerings   string `json:"offerings"`
	AboutVendor string `json:"about_vendor"`
	Licensed    bool   `json:"is_licensed"`
	StoryID     string `json:"story_id"`
}

type StallUpdate struct {
	StallName  string `json:"stall_name"`
	IsOpen     int    `json:"is_open"`
	LastActive string `json:"last_active"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

type Review struct {
	StallID       string `json:"stall_id"`
	UserID        string `json:"user_id"`
	ReviewContent string `json:"review_content"`
}

type Repository interface {
	CreateStall(ctx context.Context, reqBody Stall) error
	UpdateStall(ctx context.Context, stallId string, stall Stall) error
	RemoveStall(ctx context.Context, reqBody Stall) error
	GetStallFromStallID(ctx context.Context, stallID string) (Stall, error)
	GetUsersAroundStall(ctx context.Context, latitude float32, longitude float32) ([]domain_user.User, error)
}

type Usecase interface {
	CreateStall(ctx context.Context, reqBody Stall, assetArr []string) (string, error)
	CreateStallReview(ctx context.Context, reviewBody Review) error
	UpdateStall(ctx context.Context, stallID string, reqStall StallUpdate) (Stall, error)
}
