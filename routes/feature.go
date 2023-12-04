package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func FeatureRoutes(e *echo.Group) {
	featureRepository := repositories.RepositoryFeature(mysql.DB)
	h := handlers.HandlerFeature(featureRepository)

	e.GET("/features", h.FindFeatures)
	e.GET("/featyre/:id", h.GetFeature)
	e.POST("/featyre", middleware.Auth(h.CreateFeature))
	e.PATCH("/featyre/:id", middleware.Auth(h.UpdateFeature))
	e.DELETE("/featyre/:id", middleware.Auth(h.DeleteFeature))
}
