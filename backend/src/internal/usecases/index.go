package usecases

import (
	domain_user "htf/src/internal/domain/user"
	repositories "htf/src/internal/repository"
	usecase_user "htf/src/internal/usecases/user"
	"htf/src/utils"

	"gorm.io/gorm"
)

type UseCases struct {
	User domain_user.UseCase
}

func InitUseCases(config *utils.Config, db *gorm.DB, repos repositories.Repositories) *UseCases {
	userUseCase := usecase_user.NewUserUsecase(config, db, repos.UserRepo)
	return &UseCases{
		User: userUseCase,
	}
}
