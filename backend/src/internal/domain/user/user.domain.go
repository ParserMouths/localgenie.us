package domain_user

import "context"

type User struct {
	UserID       string `json:"user_id" gorm:"primaryKey"`
	IsVendor     string `json:"is_vendor"`
	Username     string `json:"username"`
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	PhoneNumber  int    `json:"phone_no"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Subscription string `json:"subscription"`
	StoryID      string `json:"story_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

type LoginUser struct {
	Email    string
	Username string
	Password string
}

type TokenReturn struct {
	Token   string `json:"token"`
	UserID  string `json:"user_id"`
	StallID string `json:"stall_id"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserFromUsername(ctx context.Context, username string) User
	GetStallIdFromUserId(ctx context.Context, userID string) string
}

type UseCase interface {
	CreateUser(ctx context.Context, userBody User) (User, error)
	VerifyUser(ctx context.Context, loginUser LoginUser) (bool, string, TokenReturn)
	GenerateAuthToken(ctx context.Context, loginUser LoginUser) (string, error)
}
