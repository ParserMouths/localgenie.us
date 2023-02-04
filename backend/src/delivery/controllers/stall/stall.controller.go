package controller_stall

import (
	"encoding/json"
	"fmt"
	domain_stall "htf/src/internal/domain/stall"
	"htf/src/utils"
	storyblok "htf/src/utils/storyblok"
	"path/filepath"

	// "time"

	"github.com/gofiber/fiber/v2"
)

type StallController interface {
	StallTest(fiberHandler *fiber.Ctx) (err error)
	CreateStall(fiberHandler *fiber.Ctx) (err error)
	UpdateStall(fiberHandler *fiber.Ctx) (err error)
	RemoveStall(fiberHandler *fiber.Ctx) (err error)
	QueryStall(fiberHandler *fiber.Ctx) (err error)
	CreateReview(fiberHandler *fiber.Ctx) (err error)
}

type controller struct {
	config *utils.Config
	stall  domain_stall.Usecase
}

func (c *controller) StallTest(fiberHandler *fiber.Ctx) (err error) {
	return fiberHandler.SendString("Hello from user")
}

func (c *controller) QueryStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) UpdateStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	stallId := fiberHandler.Params("id")
	var reqBody domain_stall.StallUpdate
	err = json.Unmarshal(fiberHandler.Body(), &reqBody)
	if err != nil {
		fmt.Println(err)
		return fiberHandler.SendString("error marsheling")
	}
	_, err = c.stall.UpdateStall(fiberHandler.Context(), stallId, reqBody)
	if err != nil {
		fmt.Println(err)
	}
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) RemoveStall(fiberHandler *fiber.Ctx) (err error) {
	//TODO
	return fiberHandler.SendString("Hello from user")
}
func (c *controller) CreateStall(fiberHandler *fiber.Ctx) (err error) {
	form, err := fiberHandler.MultipartForm()
	if err != nil {
		fmt.Println(err)
		return fiberHandler.JSON(fiberHandler.JSON(fiber.Map{
			"error": "error handling multipart form",
		}))
	}
	fmt.Println("###$##%#%# jere")
	ownerID := form.Value["owner"]
	stallName := form.Value["stall_name"]
	latitude := form.Value["latitude"]
	longitude := form.Value["longitude"]
	offering := form.Value["offering"]
	aboutVendor := form.Value["about_vendor"]

	files := form.File["files"]
	var assetArr []string
	for _, file := range files {
		fmt.Println(file.Filename)
		err := fiberHandler.SaveFile(file, fmt.Sprintf("./tmp/%s", file.Filename))
		if err != nil {
			fmt.Println(err)
			return fiberHandler.JSON(fiber.Map{
				"error": "error saving files locally",
			})
		}
		tmpPath := filepath.Join(c.config.ProjectRoot, "..", "tmp", file.Filename)
		assetLink, err := storyblok.CreateAsset(c.config, tmpPath)
		assetArr = append(assetArr, assetLink)
	}

	reqBody := &domain_stall.Stall{
		OwnerID:     ownerID[0],
		StallName:   stallName[0],
		Latitude:    latitude[0],
		Longitude:   longitude[0],
		Offerings:   offering[0],
		AboutVendor: aboutVendor[0],
	}

	sll, err := c.stall.CreateStall(fiberHandler.Context(), *reqBody, assetArr)
	if err != nil {
		fmt.Println(err)
		return fiberHandler.JSON(fiber.Map{
			"err": "error create story",
		})
	}
	return fiberHandler.JSON(fiber.Map{"stall_id": sll})
}

func (c *controller) CreateReview(fiberHandler *fiber.Ctx) (err error) {
	var reqBody domain_stall.Review
	json.Unmarshal(fiberHandler.Body(), &reqBody)
	err = c.stall.CreateStallReview(fiberHandler.Context(), reqBody)
	if err != nil {
		return fiberHandler.JSON(fiber.Map{
			"error": "error creating review",
		})
	}
	return fiberHandler.JSON(fiber.Map{
		"success": "review created",
	})
}

func NewStallController(config *utils.Config, stallUseCase domain_stall.Usecase) StallController {
	return &controller{
		config: config,
		stall:  stallUseCase,
	}
}
