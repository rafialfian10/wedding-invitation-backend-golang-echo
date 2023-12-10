package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func FooterRoutes(e *echo.Group) {
	footerRepository := repositories.RepositoryFooter(mysql.DB)
	h := handlers.HandlerFooter(footerRepository)

	e.GET("/footers", h.FindFooters)
	e.GET("/footer/:id", h.GetFooter)
	e.POST("/footer", middleware.Auth(h.CreateFooter))
	e.PATCH("/footer/:id", middleware.Auth(h.UpdateFooter))
	e.DELETE("/footer/:id", middleware.Auth(h.DeleteFooter))
}
