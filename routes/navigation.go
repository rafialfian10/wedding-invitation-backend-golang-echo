package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func NavigationRoutes(e *echo.Group) {
	navigationRepository := repositories.RepositoryNavigation(mysql.DB)
	h := handlers.HandlerNavigation(navigationRepository)

	e.GET("/navigations", h.FindNavigations)
	e.GET("/navigation/:id", h.GetNavigation)
	e.POST("/navigation", middleware.Auth(h.CreateNavigation))
	e.PATCH("/navigation/:id", middleware.Auth(h.UpdateNavigation))
	e.DELETE("/navigation/:id", middleware.Auth(h.DeleteNavigation))
}
