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

type handlerFagContent struct {
	FagContentRepository repositories.FagContentRepository
}

func HandlerFagContent(FagContentRepository repositories.FagContentRepository) *handlerFagContent {
	return &handlerFagContent{FagContentRepository}
}

func (h *handlerFagContent) FindFagContents(c echo.Context) error {
	fagContents, err := h.FagContentRepository.FindFagContents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleFagContentResponse(fagContents)})
}

func (h *handlerFagContent) GetFagContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var fagContent models.FagContent
	fagContent, err := h.FagContentRepository.GetFagContent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFagContentResponse(fagContent)})
}

func (h *handlerFagContent) CreateFagContent(c echo.Context) error {
	var err error

	var request dto.CreateFagContentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fagContent := models.FagContent{
		Question: request.Question,
		Answer:   request.Answer,
	}

	fagContent, err = h.FagContentRepository.CreateFagContent(fagContent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fagContent, _ = h.FagContentRepository.GetFagContent(fagContent.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFagContentResponse(fagContent)})
}

func (h *handlerFagContent) UpdateFagContent(c echo.Context) error {
	var err error

	request := dto.UpdateFagContentRequest{
		Question: c.FormValue("question"),
		Answer:   c.FormValue("answer"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	content, err := h.FagContentRepository.GetFagContent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Question != "" {
		content.Question = request.Question
	}

	if request.Answer != "" {
		content.Answer = request.Answer
	}

	data, err := h.FagContentRepository.UpdateFagContent(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerFagContent) DeleteFagContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	content, err := h.FagContentRepository.GetFagContent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FagContentRepository.DeleteFagContent(content, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFagContentResponse(data)})
}

func ConvertFagContentResponse(fagContent models.FagContent) models.FagContentResponse {
	var result models.FagContentResponse
	result.ID = fagContent.ID
	result.Question = fagContent.Question
	result.Answer = fagContent.Answer
	result.Option = fagContent.Option

	return result
}

func ConvertMultipleFagContentResponse(fagContents []models.FagContent) []models.FagContentResponse {
	var result []models.FagContentResponse

	for _, content := range fagContents {
		fagContentResponse := models.FagContentResponse{
			ID:       content.ID,
			Question: content.Question,
			Answer:   content.Answer,
			Option:   content.Option,
		}
		result = append(result, fagContentResponse)
	}
	return result
}
