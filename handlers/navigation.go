package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "wedding/dto"
	"wedding/models"
	"wedding/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerNavigation struct {
	NavigationRepository repositories.NavigationRepository
}

func HandlerNavigation(NavigationRepository repositories.NavigationRepository) *handlerNavigation {
	return &handlerNavigation{NavigationRepository}
}

func (h *handlerNavigation) FindNavigations(c echo.Context) error {
	navigations, err := h.NavigationRepository.FindNavigations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleNavigationResponse(navigations)})
}

func (h *handlerNavigation) GetNavigation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var navigation models.Navigation
	navigation, err := h.NavigationRepository.GetNavigation(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertNavigationResponse(navigation)})
}

func (h *handlerNavigation) CreateNavigation(c echo.Context) error {
	var err error

	// Parse form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid form data"})
	}

	var requests []dto.CreateNavigationRequest

	// Access fields from form data
	for _, value := range form.Value {
		for _, v := range value {
			var request dto.CreateNavigationRequest
			err := json.Unmarshal([]byte(v), &request)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			}
			requests = append(requests, request)
		}
	}

	var navigations []models.Navigation
	validation := validator.New()

	for _, request := range requests {
		err = validation.Struct(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		}

		navigation := models.Navigation{
			Name:        request.Name,
			Description: request.Description,
			Href:        request.Href,
		}

		navigations = append(navigations, navigation)
	}

	for i := range navigations {
		navigations[i], err = h.NavigationRepository.CreateNavigation(navigations[i])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		}
	}

	var response []models.NavigationResponse
	for _, navigation := range navigations {
		navigation, _ = h.NavigationRepository.GetNavigation(navigation.ID)
		response = append(response, ConvertNavigationResponse(navigation))
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: response})
}

func (h *handlerNavigation) UpdateNavigation(c echo.Context) error {
	var err error

	request := dto.UpdateNavigationRequest{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Href:        c.FormValue("href"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	navigation, err := h.NavigationRepository.GetNavigation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		navigation.Name = request.Name
	}

	if request.Description != "" {
		navigation.Description = request.Description
	}

	if request.Href != "" {
		navigation.Href = request.Href
	}

	data, err := h.NavigationRepository.UpdateNavigation(navigation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerNavigation) DeleteNavigation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	navigation, err := h.NavigationRepository.GetNavigation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.NavigationRepository.DeleteNavigation(navigation, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertNavigationResponse(data)})
}

func ConvertNavigationResponse(navigation models.Navigation) models.NavigationResponse {
	var result models.NavigationResponse

	result.ID = navigation.ID
	result.Name = navigation.Name
	result.Description = navigation.Description
	result.Href = navigation.Href

	return result
}

func ConvertMultipleNavigationResponse(navigations []models.Navigation) []models.NavigationResponse {
	var result []models.NavigationResponse

	for _, navigation := range navigations {
		navigationResponse := models.NavigationResponse{
			ID:          navigation.ID,
			Name:        navigation.Name,
			Description: navigation.Description,
			Href:        navigation.Href,
		}

		result = append(result, navigationResponse)
	}

	return result
}
