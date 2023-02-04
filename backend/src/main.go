package main

import (
	"context"
	"fmt"
	"htf/src/delivery"
	"htf/src/internal/repository"
	repo_user "htf/src/internal/repository/user"
	"htf/src/internal/usecases"
	"htf/src/utils"
	"htf/src/utils/database"
)

func main() {
	globalContext := context.Background()
	fmt.Println("htf!!")

	config := utils.NewConfig()
	db := database.NewDatabaseClient(globalContext, config)

	_, err := database.NewRedisClient(globalContext, config)

	if err != nil {
		fmt.Println(err)
	}
	repositories := repository.Repositories{
		UserRepo: repo_user.NewUserRepository(config, db),
	}

	useCases := usecases.InitUseCases(config, db, repositories)

	delivery.NewRestDelivery(globalContext, config, *useCases)
}
