package usecase_notification

import (
	"context"
	"encoding/json"
	"fmt"
	domain_notification "htf/src/internal/domain/notification"
	"htf/src/utils"

	webpush "github.com/SherClockHolmes/webpush-go"

	"gorm.io/gorm"
)

const (
	// subscription    = `{"endpoint":"https://fcm.googleapis.com/fcm/send/eLnUTCnCdRQ:APA91bEYHLM-9a1iUveU47d8nYHeTDkGze49LHG7gyTzrFdIcegkh_E1gyz08iuJywO8WtIN27frDq44wMowAbN2K_UPKA0s4izmrqAjkodgghGhH0WzFb9CudkUnkkjjyP7_8h7-P44","expirationTime":null,"keys":{"p256dh":"BO-TmO1EDO2H15tKqRng-asnvJvt7bkM1JNncTc3fy-mjroNKo0UAvNJur6qhPHUJsENlXdIRUpklkRazbJfl5U","auth":"LsbAoCwxdX-8LxEQ9hZbTQ"}}`
	vapidPublicKey  = "BMLTD4SXRjPwfFAWZCOcv9_IyWoMGr1FX1SLTgtMdTLkh5NJu6qODaju484eyptfd1m7IZl037nDQMXPcfMpRUE"
	vapidPrivateKey = "RPbgQypLEeJX-LKURKIeZKX_n9hSP9764Bmi-yl1AK0"
)

type notificationUsecase struct {
	config           *utils.Config
	db               *gorm.DB
	notificationRepo domain_notification.Repository
}

func (handler *notificationUsecase) CreateNewSubscription(ctx context.Context, reqBody domain_notification.Subscription) error {
	result := handler.db.Create(&reqBody)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (handler *notificationUsecase) SendNotificationUsecase(ctx context.Context, subscription string) error {
	ppl := domain_notification.PushPayload{
		Content: "content",
		Title:   "title",
	}
	fmt.Println(ppl)
	b, _ := json.Marshal(ppl)

	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	resp, err := webpush.SendNotification([]byte(b), s, &webpush.Options{
		Subscriber:      "example@example.com",
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})

	if err != nil {
		fmt.Errorf("notification.usecase.SendNotificationUsecase : %v", err)
		panic(err)
	}

	defer resp.Body.Close()
	return nil
}

func NewNotificationUsecase(config *utils.Config, db *gorm.DB, notificationRepo domain_notification.Repository) *notificationUsecase {
	return &notificationUsecase{
		config:           config,
		db:               db,
		notificationRepo: notificationRepo,
	}
}
