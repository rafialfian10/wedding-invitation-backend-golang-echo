package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func ContentRoutes(e *echo.Group) {
	contentRepository := repositories.RepositoryContent(mysql.DB)
	h := handlers.HandlerContent(contentRepository)

	e.GET("/contents", h.FindContents)
	e.GET("/content/:id", h.GetContent)
	e.POST("/content", middleware.Auth(h.CreateContent))
	e.PATCH("/content/:id", middleware.Auth(h.UpdateContent))
	e.DELETE("/content/:id", middleware.Auth(h.DeleteContent))
}
