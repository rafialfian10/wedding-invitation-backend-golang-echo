package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/register_admin", h.RegisterAdmin)
	e.POST("/login", h.Login)
	e.GET("/check_auth", middleware.Auth(h.CheckAuth)) // result all data user login
}
