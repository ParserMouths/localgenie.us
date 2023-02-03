package domain_user

import "context"

type User struct {
	UserID       string `json:"user_id" gorm:"primaryKey"`
	IsVendor     string `json:"is_vendor"`
	Username     string `json:"username"`
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	Location     string `json:"location"`
	PhoneNumber  int    `json:"phone_no"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Subscription string `json:"subscription"`
}

type LoginUser struct {
	Email    string
	Username string
	Password string
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
}

type UseCase interface {
	CreateUser(ctx context.Context, userBody User) (User, error)
}
