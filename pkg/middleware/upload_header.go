package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")

		if file != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			tempFile, err := ioutil.TempFile("uploads/header", "header-*.png")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()
			c.Set("dataHeader", data)
			return next(c)
		}

		c.Set("dataHeader", "")
		return next(c)
	}
}
