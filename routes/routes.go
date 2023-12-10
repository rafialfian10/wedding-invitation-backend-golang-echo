package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	UserRoutes(e)
	PricingRoutes(e)
	ContentRoutes(e)
	FeatureRoutes(e)
	HeaderRoutes(e)
	NavigationRoutes(e)
	FagRoutes(e)
	FagContentRoutes(e)
	OptionRoutes(e)
	FooterRoutes(e)
}
