package usecases

import (
	domain_notification "htf/src/internal/domain/notification"
	domain_user "htf/src/internal/domain/user"
	repositories "htf/src/internal/repository"
	usecase_notification "htf/src/internal/usecases/notification"
	usecase_user "htf/src/internal/usecases/user"
	"htf/src/utils"

	"gorm.io/gorm"
)

type UseCases struct {
	User         domain_user.UseCase
	Notification domain_notification.Usecase
}

func InitUseCases(config *utils.Config, db *gorm.DB, repos repositories.Repositories) *UseCases {
	userUseCase := usecase_user.NewUserUsecase(config, db, repos.UserRepo)
	notificationUseCase := usecase_notification.NewNotificationUsecase(config, db, repos.NotificationRepo)
	return &UseCases{
		User:         userUseCase,
		Notification: notificationUseCase,
	}
}
