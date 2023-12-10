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

type handlerContent struct {
	ContentRepository repositories.ContentRepository
}

func HandlerContent(ContentRepository repositories.ContentRepository) *handlerContent {
	return &handlerContent{ContentRepository}
}

func (h *handlerContent) FindContents(c echo.Context) error {
	contents, err := h.ContentRepository.FindContents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleContentResponse(contents)})
}

func (h *handlerContent) GetContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var content models.Content
	content, err := h.ContentRepository.GetContent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertContentResponse(content)})
}

func (h *handlerContent) CreateContent(c echo.Context) error {
	var err error

	price, _ := strconv.Atoi(c.FormValue("price"))
	mostPopuler, _ := strconv.ParseBool(c.FormValue("most_populer"))
	custom, _ := strconv.ParseBool(c.FormValue("custom"))

	request := dto.CreateContentRequest{
		Name:        c.FormValue("name"),
		Href:        c.FormValue("href"),
		Price:       price,
		Description: c.FormValue("description"),
		MostPopuler: mostPopuler,
		Custom:      custom,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	content := models.Content{
		Name:        request.Name,
		Href:        request.Href,
		Price:       request.Price,
		Description: request.Description,
		MostPopuler: false,
		Custom:      false,
	}

	content, err = h.ContentRepository.CreateContent(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	content, _ = h.ContentRepository.GetContent(content.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertContentResponse(content)})
}

func (h *handlerContent) UpdateContent(c echo.Context) error {
	var err error

	price, _ := strconv.Atoi(c.FormValue("price"))
	mostPopuler, _ := strconv.ParseBool(c.FormValue("most_populer"))
	custom, _ := strconv.ParseBool(c.FormValue("custom"))

	request := dto.UpdateContentRequest{
		Name:        c.FormValue("name"),
		Href:        c.FormValue("href"),
		Price:       price,
		Description: c.FormValue("description"),
		MostPopuler: mostPopuler,
		Custom:      custom,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	content, err := h.ContentRepository.GetContent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		content.Name = request.Name
	}

	if request.Href != "" {
		content.Href = request.Href
	}

	if request.Price != 0 {
		content.Price = request.Price
	}

	if request.Description != "" {
		content.Description = request.Description
	}

	content.MostPopuler = request.MostPopuler
	content.Custom = request.Custom

	data, err := h.ContentRepository.UpdateContent(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerContent) DeleteContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	content, err := h.ContentRepository.GetContent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ContentRepository.DeleteContent(content, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertContentResponse(data)})
}

func ConvertContentResponse(content models.Content) models.ContentResponse {
	var result models.ContentResponse
	result.ID = content.ID
	result.Name = content.Name
	result.Href = content.Href
	result.Price = content.Price
	result.Description = content.Description
	result.MostPopuler = content.MostPopuler
	result.Custom = content.Custom
	result.Feature = content.Feature

	return result
}

func ConvertMultipleContentResponse(contents []models.Content) []models.ContentResponse {
	var result []models.ContentResponse

	for _, content := range contents {
		contentResponse := models.ContentResponse{
			ID:          content.ID,
			Name:        content.Name,
			Href:        content.Href,
			Price:       content.Price,
			Description: content.Description,
			MostPopuler: content.MostPopuler,
			Custom:      content.Custom,
			Feature:     content.Feature,
		}
		result = append(result, contentResponse)
	}
	return result
}
