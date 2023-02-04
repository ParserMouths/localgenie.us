package usecase_stall

import (
	"context"
	"encoding/json"
	"fmt"
	domain_notification "htf/src/internal/domain/notification"
	domain_stall "htf/src/internal/domain/stall"
	domain_storyblok "htf/src/internal/domain/storyblok"
	domain_user "htf/src/internal/domain/user"
	"htf/src/utils"
	notif "htf/src/utils/notification"
	"htf/src/utils/storyblok"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stallUsecase struct {
	config              *utils.Config
	db                  *gorm.DB
	stallRepo           domain_stall.Repository
	notificationUsecase domain_notification.Usecase
}

func NewStallUsecase(config *utils.Config, db *gorm.DB, stallRepo domain_stall.Repository, notifUsecase domain_notification.Usecase) *stallUsecase {
	return &stallUsecase{
		config:    config,
		db:        db,
		stallRepo: stallRepo,
	}
}

func (handler *stallUsecase) CreateStall(ctx context.Context, reqStall domain_stall.Stall, assetArr []string) (domain_stall.Stall, error) {
	stallID := uuid.New().String()
	timeStr := time.Now().String()

	newStall := &domain_stall.Stall{
		StallID:     stallID,
		OwnerID:     reqStall.OwnerID,
		StallName:   reqStall.StallName,
		IsOpen:      1,
		CreatedAt:   timeStr,
		Rating:      "0",
		LastActive:  timeStr,
		Latitude:    reqStall.Latitude,
		Longitude:   reqStall.Longitude,
		Offerings:   reqStall.Offerings,
		AboutVendor: reqStall.AboutVendor,
	}

	latFloat, err := strconv.ParseFloat(newStall.Latitude, 32)
	longFloat, err := strconv.ParseFloat(newStall.Longitude, 32)
	if err != nil {
		fmt.Println(err)
		return domain_stall.Stall{}, err
	}

	newStory := &domain_storyblok.StoryPayload{
		Story: &domain_storyblok.StoryStruct{
			Name: newStall.StallName,
			Slug: newStall.StallID,
			Content: &domain_storyblok.StoryContent{
				Component:   "test",
				StallID:     stallID,
				OwnerID:     newStall.OwnerID,
				StallName:   newStall.StallName,
				IsOpen:      newStall.IsOpen,
				CreatedAt:   newStall.CreatedAt,
				Rating:      newStall.Rating,
				Offering:    newStall.Offerings,
				AboutVendor: newStall.AboutVendor,
				Latitude:    float32(latFloat),
				Longitude:   float32(longFloat),
				Assets:      assetArr,
				LastActive:  newStall.LastActive,
			},
		},
	}
	b, _ := json.Marshal(newStory)
	fmt.Println(string(b))

	storyID, err := storyblok.CreateStore(handler.config, newStory)
	fmt.Println(err)
	newStall.StoryID = storyID
	handler.stallRepo.CreateStall(ctx, *newStall)
	// also manage creating stall in storyblok using management API
	return *newStall, nil
}

func (handler *stallUsecase) UpdateStall(ctx context.Context, stallID string, reqStall domain_stall.StallUpdate) (domain_stall.Stall, error) {
	timeStr := time.Now().String()

	stall, err := handler.stallRepo.GetStallFromStallID(ctx, stallID)
	if err != nil {
		return domain_stall.Stall{}, err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://mapi.storyblok.com/v1/spaces/%s/stories/%s", handler.config.SpaceID, stall.StoryID), nil)
	req.Header.Set("Authorization", handler.config.StoryBlokOAuth)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Mauu", err)
	}
	defer res.Body.Close()
	var currentStory domain_storyblok.StoryPayload
	bodyByte, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(bodyByte, &currentStory)
	if err != nil {
		fmt.Println("mulu mulu", err)
	}
	intStoryID, _ := strconv.Atoi(stall.StoryID)
	newStall := &domain_stall.Stall{
		StallID:     stall.StallID,
		OwnerID:     stall.OwnerID,
		StallName:   stall.StallName,
		IsOpen:      reqStall.IsOpen,
		CreatedAt:   stall.CreatedAt,
		Rating:      stall.Rating,
		LastActive:  timeStr,
		Latitude:    reqStall.Latitude,
		Longitude:   reqStall.Longitude,
		Offerings:   stall.Offerings,
		AboutVendor: stall.AboutVendor,
		Licensed:    stall.Licensed,
		StoryID:     stall.StoryID,
	}

	var latFloat, longFloat float64
	newStoryContent := &domain_storyblok.StoryContent{
		Component:   "test",
		Assets:      currentStory.Story.Content.Assets,
		StallID:     currentStory.Story.Content.StallID,
		OwnerID:     currentStory.Story.Content.OwnerID,
		IsOpen:      newStall.IsOpen,
		LastActive:  newStall.LastActive,
		Rating:      currentStory.Story.Content.Rating,
		Offering:    currentStory.Story.Content.Offering,
		AboutVendor: currentStory.Story.Content.AboutVendor,
		CreatedAt:   currentStory.Story.Content.CreatedAt,
		StallName:   currentStory.Story.Content.StallName,
	}

	if reqStall.Latitude != "" && reqStall.Longitude != "" {
		latFloat, err = strconv.ParseFloat(newStall.Latitude, 32)
		longFloat, err = strconv.ParseFloat(newStall.Longitude, 32)
		newStoryContent.Latitude = float32(latFloat)
		newStoryContent.Longitude = float32(longFloat)
	}
	if newStall.IsOpen == 1 {
		usersNearby, _ := handler.stallRepo.GetUsersAroundStall(ctx, float32(latFloat), float32(longFloat))
		for _, user := range usersNearby {
			go func(user domain_user.User) {
				notif.SendNotif(handler.config, user.Subscription, "Your favourite stall is open.", fmt.Sprintf("%s is open now.", stall.StallName))
			}(user)
		}
	}
	if err != nil {
		fmt.Println("JHeereeer", err)
		return domain_stall.Stall{}, err
	}

	newStory := &domain_storyblok.StoryPayload{
		Story: &domain_storyblok.StoryStruct{
			Name:    newStall.StallName,
			Slug:    newStall.StallID,
			StoryID: intStoryID,
			Content: newStoryContent,
		},
	}

	_, err = storyblok.UpdateStory(handler.config, stall.StoryID, newStory)
	fmt.Println("Heereeee", err)

	handler.stallRepo.UpdateStall(ctx, stallID, *newStall)
	// also manage creating stall in storyblok using management API
	return *newStall, nil
}

func (handler *stallUsecase) CreateStallReview(ctx context.Context, reviewBody domain_stall.Review) error {
	results := handler.db.Create(reviewBody)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
