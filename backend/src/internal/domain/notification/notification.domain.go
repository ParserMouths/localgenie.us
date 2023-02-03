package domain_notification

import "context"

type Subscription struct {
	SubscriptionID string `json:"subsrciption"`
	UserID         string `json:"user_id"`
}

type PushPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Repository interface {
}

type Usecase interface {
	CreateNewSubscription(ctx context.Context, reqBody Subscription) error
	SendNotificationUsecase(ctx context.Context, reqBody PushPayload) error
}
