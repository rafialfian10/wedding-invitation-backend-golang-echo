package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func FagContentRoutes(e *echo.Group) {
	fagContentRepository := repositories.RepositoryFagContent(mysql.DB)
	h := handlers.HandlerFagContent(fagContentRepository)

	e.GET("/fag_contents", h.FindFagContents)
	e.GET("/fag_content/:id", h.GetFagContent)
	e.POST("/fag_content", middleware.Auth(h.CreateFagContent))
	e.PATCH("/fag_content/:id", middleware.Auth(h.UpdateFagContent))
	e.DELETE("/fag_content/:id", middleware.Auth(h.DeleteFagContent))
}
