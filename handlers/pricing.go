package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	dto "wedding/dto"
	"wedding/models"
	"wedding/repositories"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// var path_image = "http://localhost:5000/uploads/image/"

type handlerPricing struct {
	PricingRepository repositories.PricingRepository
}

func HandlerPricing(PricingRepository repositories.PricingRepository) *handlerPricing {
	return &handlerPricing{PricingRepository}
}

func (h *handlerPricing) FindPricings(c echo.Context) error {
	pricings, err := h.PricingRepository.FindPricings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	for i, pricing := range pricings {
		pricings[i].Image = pricing.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultiplePricingResponse(pricings)})
}

func (h *handlerPricing) GetPricing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var pricing models.Pricing
	pricing, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertPricingResponse(pricing)})
}

func (h *handlerPricing) CreatePricing(c echo.Context) error {
	var err error

	var request dto.CreatePricingRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	dataContext := c.Get("dataImage")
	if dataContext == nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataImage is nil"})
	}

	filepath, ok := dataContext.(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataImage is not a string"})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "wedding"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// if image nil return ""
	var imageSecureURL string
	if resp != nil && resp.SecureURL != "" {
		imageSecureURL = resp.SecureURL
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	pricing := models.Pricing{
		Title:       request.Title,
		Caption:     request.Caption,
		Description: request.Description,
		Image:       imageSecureURL,
	}

	for _, content := range request.Contents {
		contentData := models.ContentResponse{
			Name:        content.Name,
			Href:        content.Href,
			Price:       content.Price,
			Description: content.Description,
			MostPopuler: false,
			Custom:      false,
		}

		for _, feature := range content.Features {
			featureData := models.FeatureResponse{
				Feature: feature.Feature,
			}
			contentData.Feature = append(contentData.Feature, featureData)
		}

		pricing.Content = append(pricing.Content, contentData)
	}

	pricing, err = h.PricingRepository.CreatePricing(pricing)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	pricing, _ = h.PricingRepository.GetPricing(pricing.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: pricing})
}

func (h *handlerPricing) UpdatePricing(c echo.Context) error {
	var err error

	dataContext := c.Get("dataImage")
	if dataContext == nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataImage is nil"})
	}

	filepath, ok := dataContext.(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataImage is not a string"})
	}

	// cloudinary
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "wedding"})

	if err != nil {
		fmt.Println(err.Error())
	}

	var imageSecureURL string
	if resp != nil && resp.SecureURL != "" {
		imageSecureURL = resp.SecureURL
	}

	request := dto.UpdatePricingRequest{
		Caption:     c.FormValue("caption"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Image:       imageSecureURL,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	pricing, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Caption != "" {
		pricing.Caption = request.Caption
	}

	if request.Title != "" {
		pricing.Title = request.Title
	}

	if request.Description != "" {
		pricing.Description = request.Description
	}

	if request.Image != "" {
		pricing.Image = request.Image
	}

	data, err := h.PricingRepository.UpdatePricing(pricing)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerPricing) DeletePricing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.PricingRepository.DeletePricing(movie, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertPricingResponse(data)})
}

func (h *handlerPricing) DeleteImage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.PricingRepository.DeleteImage(id); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	pricing, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertPricingResponse(pricing)})
}

func ConvertPricingResponse(pricing models.Pricing) models.PricingResponse {
	var result models.PricingResponse
	result.ID = pricing.ID
	result.Caption = pricing.Caption
	result.Title = pricing.Title
	result.Description = pricing.Description
	result.Image = pricing.Image
	result.Content = pricing.Content
	// result.UserID = pricing.UserID
	// result.User = pricing.User

	return result
}

func ConvertMultiplePricingResponse(pricings []models.Pricing) []models.PricingResponse {
	var result []models.PricingResponse

	for _, pricing := range pricings {
		pricingResponse := models.PricingResponse{
			ID:          pricing.ID,
			Caption:     pricing.Caption,
			Title:       pricing.Title,
			Description: pricing.Description,
			Image:       pricing.Image,
			Content:     pricing.Content,
			// UserID:      pricing.UserID,
			// User:        pricing.User,
		}
		result = append(result, pricingResponse)
	}
	return result
}
