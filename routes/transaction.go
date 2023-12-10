package routes

import (
	"wedding/handlers"
	"wedding/pkg/middleware"
	"wedding/pkg/mysql"
	"wedding/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transactions-by-user", middleware.Auth(h.GetAllTransactionByUser))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.POST("/notification", h.Notification)
	e.PATCH("/transaction/:id", middleware.Auth(h.UpdateTransaction))
	e.PATCH("/transaction-admin/:id", middleware.Auth(h.UpdateTransactionByAdmin))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
}
