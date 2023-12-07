package handlers

import (
	"net/http"
	"strconv"
	dto "wedding/dto"
	"wedding/models"
	"wedding/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerFeature struct {
	FeatureRepository repositories.FeatureRepository
}

func HandlerFeature(FeatureRepository repositories.FeatureRepository) *handlerFeature {
	return &handlerFeature{FeatureRepository}
}

func (h *handlerFeature) FindFeatures(c echo.Context) error {
	features, err := h.FeatureRepository.FindFeatures()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleFeatureResponse(features)})
}

func (h *handlerFeature) GetFeature(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var feature models.Feature
	feature, err := h.FeatureRepository.GetFeature(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFeatureResponse(feature)})
}

func (h *handlerFeature) CreateFeature(c echo.Context) error {
	var err error

	var request dto.CreateFeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	feature := models.Feature{
		Feature: request.Feature,
	}

	feature, err = h.FeatureRepository.CreateFeature(feature)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	feature, _ = h.FeatureRepository.GetFeature(feature.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFeatureResponse(feature)})
}

func (h *handlerFeature) UpdateFeature(c echo.Context) error {
	var err error

	request := dto.UpdateFeatureRequest{
		Feature: c.FormValue("feature"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	feature, err := h.FeatureRepository.GetFeature(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Feature != "" {
		feature.Feature = request.Feature
	}

	data, err := h.FeatureRepository.UpdateFeature(feature)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerFeature) DeleteFeature(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	content, err := h.FeatureRepository.GetFeature(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FeatureRepository.DeleteFeature(content, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFeatureResponse(data)})
}

func ConvertFeatureResponse(feature models.Feature) models.FeatureResponse {
	var result models.FeatureResponse
	result.ID = feature.ID
	result.Feature = feature.Feature

	return result
}

func ConvertMultipleFeatureResponse(features []models.Feature) []models.FeatureResponse {
	var result []models.FeatureResponse

	for _, feature := range features {
		featureResponse := models.FeatureResponse{
			ID:      feature.ID,
			Feature: feature.Feature,
		}

		result = append(result, featureResponse)
	}

	return result
}
