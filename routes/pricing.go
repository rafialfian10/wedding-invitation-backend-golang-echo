package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func PricingRoutes(e *echo.Group) {
	pricingRepository := repositories.RepositoryPricing(mysql.DB)
	h := handlers.HandlerPricing(pricingRepository)

	e.GET("/pricings", h.FindPricings)
	e.GET("/pricing/:id", h.GetPricing)
	e.POST("/pricing", middleware.Auth(middleware.UploadImage(h.CreatePricing)))
	e.PATCH("/pricing/:id", middleware.Auth(middleware.UploadImage(h.UpdatePricing)))
	e.DELETE("/pricing/:id", middleware.Auth(h.DeletePricing))
}
