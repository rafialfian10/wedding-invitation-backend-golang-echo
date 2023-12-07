package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", middleware.Auth(h.GetUser))
	e.POST("/user", h.CreateUser)
	e.PATCH("/user/:id", middleware.Auth(middleware.UploadPhoto(h.UpdateUser)))
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
	e.GET("/user", middleware.Auth(h.GetProfile))
	e.DELETE("/user/:id/photo", middleware.Auth(h.DeletePhoto))
}
