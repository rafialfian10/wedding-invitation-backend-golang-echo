package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "wedding/dto"
	"wedding/models"
	"wedding/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var path_image = "http://localhost:5000/uploads/image/"

type handlerPricing struct {
	PricingRepository repositories.PricingRepository
}

func HandlerPricing(PricingRepository repositories.PricingRepository) *handlerPricing {
	return &handlerPricing{PricingRepository}
}

func (h *handlerPricing) FindPricings(c echo.Context) error {
	contents, err := h.PricingRepository.FindPricings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	for i, content := range contents {
		contents[i].Image = path_image + content.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultiplePricingResponse(contents)})
}

func (h *handlerPricing) GetPricing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var pricing models.Pricing
	pricing, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertPricingResponse(pricing)})
}

func (h *handlerPricing) CreatePricing(c echo.Context) error {
	var err error
	dataImage := c.Get("dataImage").(string)
	// fmt.Println("this is data file", dataImage)

	contentIdString := c.FormValue("content_id")

	if contentIdString == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Error: content_id form value is missing."})
	}

	var categoriesId []int
	err = json.Unmarshal([]byte(contentIdString), &categoriesId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if len(categoriesId) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Error: content_id form value is missing."})
	}

	request := dto.CreatePricingRequest{
		Caption:     c.FormValue("caption"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		ContentID:   categoriesId,
		Image:       dataImage,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	pricing := models.Pricing{
		Title:   request.Title,
		Caption: request.Caption,
		// Content:     categories,
		Description: request.Description,
		Image:       request.Image,
		UserID:      int(userId),
	}

	pricing, err = h.PricingRepository.CreatePricing(pricing)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	pricing, _ = h.PricingRepository.GetPricing(pricing.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertPricingResponse(pricing)})
}

func (h *handlerPricing) UpdatePricing(c echo.Context) error {
	var err error
	dataImage := c.Get("dataImage").(string)

	var contentId []int
	contentIdString := c.FormValue("content_id")
	err = json.Unmarshal([]byte(contentIdString), &contentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	request := dto.UpdatePricingRequest{
		Caption:     c.FormValue("caption"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Image:       dataImage,
		ContentID:   contentId,
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

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertPricingResponse(data)})
}

func (h *handlerPricing) DeleteImage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.PricingRepository.DeleteImageByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	movie, err := h.PricingRepository.GetPricing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertPricingResponse(movie)})
}

func convertPricingResponse(pricing models.Pricing) models.PricingResponse {
	var result models.PricingResponse
	result.ID = pricing.ID
	result.Caption = pricing.Caption
	result.Title = pricing.Title
	result.Description = pricing.Description
	result.Image = pricing.Image
	result.ContentId = pricing.ContentId
	result.Content = pricing.Content
	result.UserID = pricing.UserID
	result.User = pricing.User

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
			ContentId:   pricing.ContentId,
			Content:     pricing.Content,
			UserID:      pricing.UserID,
			User:        pricing.User,
		}

		result = append(result, pricingResponse)
	}

	return result
}
