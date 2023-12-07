package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func HeaderRoutes(e *echo.Group) {
	headerRepository := repositories.RepositoryHeader(mysql.DB)
	h := handlers.HandlerHeader(headerRepository)

	e.GET("/headers", h.FindHeaders)
	e.GET("/header/:id", h.GetHeader)
	e.POST("/header", middleware.Auth(middleware.UploadHeader(h.CreateHeader)))
	e.PATCH("/header/:id", middleware.Auth(middleware.UploadHeader(h.UpdateHeader)))
	e.DELETE("/header/:id", middleware.Auth(h.DeleteHeader))
	e.DELETE("/header/:id/image", middleware.Auth(h.DeleteHeaderImage))
}
