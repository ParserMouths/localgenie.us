package notification

import (
	"encoding/json"
	"fmt"
	domain_notification "htf/src/internal/domain/notification"
	"htf/src/utils"

	"github.com/SherClockHolmes/webpush-go"
)

const vapidPublicKey = "BMLTD4SXRjPwfFAWZCOcv9_IyWoMGr1FX1SLTgtMdTLkh5NJu6qODaju484eyptfd1m7IZl037nDQMXPcfMpRUE"

func SendNotif(config *utils.Config, subscription string, title string, content string) {

	ppl := domain_notification.PushPayload{
		Content: content,
		Title:   title,
	}
	b, _ := json.Marshal(ppl)

	s := &webpush.Subscription{}
	err := json.Unmarshal([]byte(subscription), s)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := webpush.SendNotification([]byte(b), s, &webpush.Options{
		Subscriber:      "example@example.com",
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: config.NotifPrivateKey,
		TTL:             30,
	})

	if err != nil {
		fmt.Errorf("notification.usecase.SendNotificationUsecase : %v", err)
		panic(err)
	}

	defer resp.Body.Close()
}
