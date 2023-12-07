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

// var path_image = "http://localhost:5000/uploads/header/"

type handlerHeader struct {
	HeaderRepository repositories.HeaderRepository
}

func HandlerHeader(handlerRepository repositories.HeaderRepository) *handlerHeader {
	return &handlerHeader{handlerRepository}
}

func (h *handlerHeader) FindHeaders(c echo.Context) error {
	headers, err := h.HeaderRepository.FindHeaders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	for i, header := range headers {
		headers[i].Image = header.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleHeaderResponse(headers)})
}

func (h *handlerHeader) GetHeader(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var header models.Header
	header, err := h.HeaderRepository.GetHeader(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertHeaderResponse(header)})
}

func (h *handlerHeader) CreateHeader(c echo.Context) error {
	var err error

	var request dto.CreateHeaderRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid request format"})
	}

	dataContext := c.Get("dataHeader")
	if dataContext == nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataHeader is nil"})
	}

	filepath, ok := dataContext.(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataHeader is not a string"})
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

	header := models.Header{
		Header:    request.Header,
		SubHeader: request.SubHeader,
		Button:    request.Button,
		Image:     imageSecureURL,
	}

	header, err = h.HeaderRepository.CreateHeader(header)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	header, _ = h.HeaderRepository.GetHeader(header.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertHeaderResponse(header)})
}

func (h *handlerHeader) UpdateHeader(c echo.Context) error {
	var err error

	dataContext := c.Get("dataHeader")
	if dataContext == nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataHeader is nil"})
	}

	filepath, ok := dataContext.(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "dataHeader is not a string"})
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

	request := dto.UpdateHeaderRequest{
		Header:    c.FormValue("header"),
		SubHeader: c.FormValue("sub_header"),
		Button:    c.FormValue("button"),
		Image:     imageSecureURL,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	header, err := h.HeaderRepository.GetHeader(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Header != "" {
		header.Header = request.Header
	}

	if request.SubHeader != "" {
		header.SubHeader = request.SubHeader
	}

	if request.Button != "" {
		header.Button = request.Button
	}

	if request.Image != "" {
		header.Image = request.Image
	}

	data, err := h.HeaderRepository.UpdateHeader(header)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerHeader) DeleteHeader(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	header, err := h.HeaderRepository.GetHeader(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.HeaderRepository.DeleteHeader(header, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertHeaderResponse(data)})
}

func (h *handlerHeader) DeleteHeaderImage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.HeaderRepository.DeleteHeaderImage(id); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	header, err := h.HeaderRepository.GetHeader(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertHeaderResponse(header)})
}

func ConvertHeaderResponse(header models.Header) models.HeaderResponse {
	var result models.HeaderResponse
	result.ID = header.ID
	result.Header = header.Header
	result.SubHeader = header.SubHeader
	result.Button = header.Button
	result.Image = header.Image

	return result
}

func ConvertMultipleHeaderResponse(headers []models.Header) []models.HeaderResponse {
	var result []models.HeaderResponse

	for _, pricing := range headers {
		headerResponse := models.HeaderResponse{
			ID:        pricing.ID,
			Header:    pricing.Header,
			SubHeader: pricing.SubHeader,
			Button:    pricing.Button,
			Image:     pricing.Image,
		}
		result = append(result, headerResponse)
	}
	return result
}
