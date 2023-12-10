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

type handlerFag struct {
	FagRepository repositories.FagRepository
}

func HandlerFag(FagRepository repositories.FagRepository) *handlerFag {
	return &handlerFag{FagRepository}
}

func (h *handlerFag) FindFags(c echo.Context) error {
	fags, err := h.FagRepository.FindFags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleFagResponse(fags)})
}

func (h *handlerFag) GetFag(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var fag models.Fag
	fag, err := h.FagRepository.GetFag(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFagResponse(fag)})
}

func (h *handlerFag) CreateFag(c echo.Context) error {
	var err error

	var request dto.CreateFagRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fag := models.Fag{
		Title:       request.Title,
		Caption:     request.Caption,
		Description: request.Description,
	}

	for _, fagContent := range request.FagContents {
		fagContentData := models.FagContentResponse{
			Question: fagContent.Question,
			Answer:   fagContent.Answer,
		}

		for _, option := range fagContent.Options {
			optionData := models.OptionResponse{
				Option: option.Option,
			}
			fagContentData.Option = append(fagContentData.Option, optionData)
		}

		fag.FagContent = append(fag.FagContent, fagContentData)
	}

	fag, err = h.FagRepository.CreateFag(fag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fag, _ = h.FagRepository.GetFag(fag.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: fag})
}

func (h *handlerFag) UpdateFag(c echo.Context) error {
	var err error

	request := dto.UpdateFagRequest{
		Caption:     c.FormValue("caption"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	fag, err := h.FagRepository.GetFag(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Caption != "" {
		fag.Caption = request.Caption
	}

	if request.Title != "" {
		fag.Title = request.Title
	}

	if request.Description != "" {
		fag.Description = request.Description
	}

	data, err := h.FagRepository.UpdateFag(fag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerFag) DeleteFag(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fag, err := h.FagRepository.GetFag(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FagRepository.DeleteFag(fag, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFagResponse(data)})
}

func ConvertFagResponse(fag models.Fag) models.FagResponse {
	var result models.FagResponse
	result.ID = fag.ID
	result.Caption = fag.Caption
	result.Title = fag.Title
	result.Description = fag.Description

	return result
}

func ConvertMultipleFagResponse(fags []models.Fag) []models.FagResponse {
	var result []models.FagResponse

	for _, fag := range fags {
		fagResponse := models.FagResponse{
			ID:          fag.ID,
			Caption:     fag.Caption,
			Title:       fag.Title,
			Description: fag.Description,
		}
		result = append(result, fagResponse)
	}
	return result
}
