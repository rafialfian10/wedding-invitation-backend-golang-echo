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

type handlerFooter struct {
	FooterRepository repositories.FooterRepository
}

func HandlerFooter(FooterRepository repositories.FooterRepository) *handlerFooter {
	return &handlerFooter{FooterRepository}
}

func (h *handlerFooter) FindFooters(c echo.Context) error {
	footers, err := h.FooterRepository.FindFooters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleFooterResponse(footers)})
}

func (h *handlerFooter) GetFooter(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var footer models.Footer
	footer, err := h.FooterRepository.GetFooter(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFooterResponse(footer)})
}

func (h *handlerFooter) CreateFooter(c echo.Context) error {
	var err error

	var request dto.CreateFooterRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	footer := models.Footer{
		FooterOneHeader:    request.FooterOneHeader,
		FooterOneContent:   request.FooterOneContent,
		FooterTwoHeader:    request.FooterTwoHeader,
		FooterTwoContent:   request.FooterTwoContent,
		FooterThreeHeader:  request.FooterThreeHeader,
		FooterThreeContent: request.FooterThreeContent,
		FooterFourHeader:   request.FooterFourHeader,
		FooterFourContent:  request.FooterFourContent,
		FooterFiveHeader:   request.FooterFiveHeader,
		FooterFiveContent:  request.FooterFiveContent,
		Copyright:          request.Copyright,
	}

	// footer.FooterTwoContent = append(footer.FooterTwoContent, request.FooterTwoContent...)
	// footer.FooterThreeContent = append(footer.FooterThreeContent, request.FooterThreeContent...)
	// footer.FooterFourContent = append(footer.FooterFourContent, request.FooterFourContent...)
	// footer.FooterFiveContent = append(footer.FooterFiveContent, request.FooterFiveContent...)

	footer, err = h.FooterRepository.CreateFooter(footer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	footer, _ = h.FooterRepository.GetFooter(footer.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: footer})
}

func (h *handlerFooter) UpdateFooter(c echo.Context) error {
	var err error

	request := dto.UpdateFooterRequest{
		FooterOneHeader:    c.FormValue("footer_one_header"),
		FooterOneContent:   c.FormValue("footer_one_content"),
		FooterTwoHeader:    c.FormValue("footer_two_header"),
		FooterTwoContent:   c.FormValue("footer_two_content"),
		FooterThreeHeader:  c.FormValue("footer_three_header"),
		FooterThreeContent: c.FormValue("footer_three_content"),
		FooterFourHeader:   c.FormValue("footer_four_header"),
		FooterFourContent:  c.FormValue("footer_four_content"),
		FooterFiveHeader:   c.FormValue("footer_five_header"),
		FooterFiveContent:  c.FormValue("footer_five_content"),
		Copyright:          c.FormValue("copyright"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	footer, err := h.FooterRepository.GetFooter(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.FooterOneHeader != "" {
		footer.FooterOneHeader = request.FooterOneHeader
	}

	if request.FooterOneContent != "" {
		footer.FooterOneContent = request.FooterOneContent
	}

	if request.FooterTwoHeader != "" {
		footer.FooterTwoHeader = request.FooterTwoHeader
	}

	if request.FooterTwoContent != "" {
		footer.FooterTwoContent = request.FooterTwoContent
	}

	if request.FooterThreeHeader != "" {
		footer.FooterThreeHeader = request.FooterThreeHeader
	}

	if request.FooterThreeContent != "" {
		footer.FooterThreeContent = request.FooterThreeContent
	}

	if request.FooterFourHeader != "" {
		footer.FooterFourHeader = request.FooterFourHeader
	}

	if request.FooterFourContent != "" {
		footer.FooterFourContent = request.FooterFourContent
	}

	if request.FooterFiveHeader != "" {
		footer.FooterFiveHeader = request.FooterFiveHeader
	}

	if request.FooterFiveContent != "" {
		footer.FooterFiveContent = request.FooterFiveContent
	}

	// if len(request.FooterFiveContent) > 0 {
	// 	footer.FooterFiveContent = request.FooterFiveContent
	// }

	if request.Copyright != "" {
		footer.Copyright = request.Copyright
	}

	data, err := h.FooterRepository.UpdateFooter(footer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerFooter) DeleteFooter(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	footer, err := h.FooterRepository.GetFooter(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FooterRepository.DeleteFooter(footer, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertFooterResponse(data)})
}

func ConvertFooterResponse(footer models.Footer) models.FooterResponse {
	var result models.FooterResponse
	result.ID = footer.ID
	result.FooterOneHeader = footer.FooterOneHeader
	result.FooterOneContent = footer.FooterOneContent
	result.FooterTwoHeader = footer.FooterTwoHeader
	result.FooterTwoContent = footer.FooterTwoContent
	result.FooterThreeHeader = footer.FooterThreeHeader
	result.FooterThreeContent = footer.FooterThreeContent
	result.FooterFourHeader = footer.FooterFourHeader
	result.FooterFourContent = footer.FooterFourContent
	result.FooterFiveHeader = footer.FooterFiveHeader
	result.FooterFiveContent = footer.FooterFiveContent
	result.Copyright = footer.Copyright

	return result
}

func ConvertMultipleFooterResponse(footers []models.Footer) []models.FooterResponse {
	var result []models.FooterResponse

	for _, footer := range footers {
		footerResponse := models.FooterResponse{
			ID:                 footer.ID,
			FooterOneHeader:    footer.FooterOneHeader,
			FooterOneContent:   footer.FooterOneContent,
			FooterTwoHeader:    footer.FooterTwoHeader,
			FooterTwoContent:   footer.FooterTwoContent,
			FooterThreeHeader:  footer.FooterThreeHeader,
			FooterThreeContent: footer.FooterThreeContent,
			FooterFourHeader:   footer.FooterFourHeader,
			FooterFourContent:  footer.FooterFourContent,
			FooterFiveHeader:   footer.FooterFiveHeader,
			FooterFiveContent:  footer.FooterFiveContent,
			Copyright:          footer.Copyright,
		}

		result = append(result, footerResponse)
	}

	return result
}
