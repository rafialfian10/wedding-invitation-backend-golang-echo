package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func OptionRoutes(e *echo.Group) {
	optionRepository := repositories.RepositoryOption(mysql.DB)
	h := handlers.HandlerOption(optionRepository)

	e.GET("/options", h.FindOptions)
	e.GET("/option/:id", h.GetOption)
	e.POST("/option", middleware.Auth(h.CreateOption))
	e.PATCH("/option/:id", middleware.Auth(h.UpdateOption))
	e.DELETE("/option/:id", middleware.Auth(h.DeleteOption))
}
