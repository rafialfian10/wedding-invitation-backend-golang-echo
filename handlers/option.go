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

type handlerOption struct {
	OptionRepository repositories.OptionRepository
}

func HandlerOption(OptionRepository repositories.OptionRepository) *handlerOption {
	return &handlerOption{OptionRepository}
}

func (h *handlerOption) FindOptions(c echo.Context) error {
	options, err := h.OptionRepository.FindOptions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleOptionResponse(options)})
}

func (h *handlerOption) GetOption(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var option models.Option
	option, err := h.OptionRepository.GetOption(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertOptionResponse(option)})
}

func (h *handlerOption) CreateOption(c echo.Context) error {
	var err error

	var request dto.CreateOptionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	option := models.Option{
		Option: request.Option,
	}

	option, err = h.OptionRepository.CreateOption(option)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	option, _ = h.OptionRepository.GetOption(option.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertOptionResponse(option)})
}

func (h *handlerOption) UpdateOption(c echo.Context) error {
	var err error

	request := dto.UpdateOptionRequest{
		Option: c.FormValue("option"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	option, err := h.OptionRepository.GetOption(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Option != "" {
		option.Option = request.Option
	}

	data, err := h.OptionRepository.UpdateOption(option)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerOption) DeleteOption(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	option, err := h.OptionRepository.GetOption(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.OptionRepository.DeleteOption(option, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertOptionResponse(data)})
}

func ConvertOptionResponse(option models.Option) models.OptionResponse {
	var result models.OptionResponse
	result.ID = option.ID
	result.Option = option.Option

	return result
}

func ConvertMultipleOptionResponse(options []models.Option) []models.OptionResponse {
	var result []models.OptionResponse

	for _, option := range options {
		optionResponse := models.OptionResponse{
			ID:     option.ID,
			Option: option.Option,
		}

		result = append(result, optionResponse)
	}

	return result
}
