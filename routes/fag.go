package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func FagRoutes(e *echo.Group) {
	fagRepository := repositories.RepositoryFag(mysql.DB)
	h := handlers.HandlerFag(fagRepository)

	e.GET("/fags", h.FindFags)
	e.GET("/fag/:id", h.GetFag)
	e.POST("/fag", middleware.Auth(h.CreateFag))
	e.PATCH("/fag/:id", middleware.Auth(h.UpdateFag))
	e.DELETE("/fag/:id", middleware.Auth(h.DeleteFag))
}
