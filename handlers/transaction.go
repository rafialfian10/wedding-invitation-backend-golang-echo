package handlers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	dto "wedding/dto"
	"wedding/models"
	"wedding/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	pricings, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleTransactionResponse(pricings)})
}

func (h *handlerTransaction) GetAllTransactionByUser(c echo.Context) error {
	userId := c.Get("userLogin").(jwt.MapClaims)["id"].(float64)

	transaction, err := h.TransactionRepository.FindTransactionsByUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertMultipleTransactionResponse(transaction)})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertTransactionResponse(transaction)})
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	userId := c.Get("userLogin").(jwt.MapClaims)["id"].(float64)

	total, _ := strconv.Atoi(c.FormValue("total"))
	pricingId, _ := strconv.Atoi(c.FormValue("pricing_id"))

	request := dto.CreateTransactionRequest{
		Total:     total,
		PricingID: pricingId,
		UserID:    int(userId),
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	var TrxIdMatch = false
	var TrxId int
	for !TrxIdMatch {
		TrxId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(TrxId)
		if transactionData.ID == 0 {
			TrxIdMatch = true
		}
	}

	newTransaction := models.Transaction{
		ID:          TrxId,
		Total:       request.Total,
		BookingDate: time.Now().UTC(),
		Status:      "pending",
		PricingID:   request.PricingID,
		UserID:      request.UserID,
	}

	transaction, err := h.TransactionRepository.CreateTransaction(newTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	TransactionAdded, _ := h.TransactionRepository.GetTransaction(transaction.ID)

	var s = snap.Client{}
	s.New(os.Getenv("TRANSACTION_SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(TransactionAdded.ID),
			GrossAmt: int64(TransactionAdded.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: TransactionAdded.User.Username,
			Email: TransactionAdded.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	updateTransaction, _ := h.TransactionRepository.UpdateTokenTransaction(snapResp.Token, TransactionAdded.ID)
	transactionUpdated, _ := h.TransactionRepository.GetTransaction(updateTransaction.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertTransactionResponse(transactionUpdated)})
}

func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	var s = snap.Client{}
	s.New(os.Getenv("TRANSACTION_SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.User.Username,
			Email: transaction.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	transaction, _ = h.TransactionRepository.UpdateTokenTransaction(snapResp.Token, id)
	transactionUpdated, _ := h.TransactionRepository.GetTransaction(id)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertTransactionResponse(transactionUpdated)})
}

func (h *handlerTransaction) UpdateTransactionByAdmin(c echo.Context) error {
	var err error

	request := dto.UpdateTransactionRequest{
		Status: c.FormValue("status"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(request.Status, transaction.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	order_id, _ := strconv.Atoi(orderId)

	fmt.Print("ini payloadnya", notificationPayload)
	fmt.Println("order id", order_id)

	transaction, _ := h.TransactionRepository.GetTransaction(order_id)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.TransactionRepository.UpdateTransaction("pending", order_id)
		} else if fraudStatus == "accept" {
			h.TransactionRepository.UpdateTransaction("success", order_id)
			SendEmail("Transaction Success", transaction)
		}
	} else if transactionStatus == "settlement" {
		h.TransactionRepository.UpdateTransaction("success", order_id)
		SendEmail("Transaction Success", transaction)
	} else if transactionStatus == "deny" {
		h.TransactionRepository.UpdateTransaction("failed", order_id)
		SendEmail("Transaction Failed", transaction)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.TransactionRepository.UpdateTransaction("failed", order_id)
		SendEmail("Transaction Failed", transaction)
	} else if transactionStatus == "pending" {
		h.TransactionRepository.UpdateTransaction("pending", order_id)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: notificationPayload})
}

func SendEmail(status string, transaction models.Transaction) {
	var CONFIG_SMTP_HOST = "smtp.gmail.com"
	var CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = "dewetour <rafialfian770@gmail.com>"
	var CONFIG_AUTH_EMAIL = os.Getenv("SYSTEM_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("SYSTEM_PASSWORD")

	var pricingTitle = transaction.Pricing.Title
	var price = strconv.Itoa(transaction.Total)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", transaction.User.Email)
	mailer.SetHeader("Subject", "Status Transaction")
	mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total payment: Rp.%s</li>
        <li>Status : %s</li>
		<li>Iklan : %s</li>
      </ul>
      </body>
    </html>`, pricingTitle, price, status, "Terima kasih"))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertTransactionResponse(data)})
}

func ConvertTransactionResponse(transaction models.Transaction) models.TransactionResponse {
	result := models.TransactionResponse{
		ID:          transaction.ID,
		Total:       transaction.Total,
		Status:      transaction.Status,
		Token:       transaction.Token,
		BookingDate: transaction.BookingDate,
		UserID:      transaction.UserID,
		User:        transaction.User,
		PricingID:   transaction.PricingID,
		Pricing: models.PricingResponse{
			ID:          transaction.Pricing.ID,
			Caption:     transaction.Pricing.Caption,
			Title:       transaction.Pricing.Title,
			Description: transaction.Pricing.Description,
			Image:       transaction.Pricing.Image,
			Content:     transaction.Pricing.Content,
		},
	}
	return result
}

func ConvertMultipleTransactionResponse(transactions []models.Transaction) []models.TransactionResponse {
	var result []models.TransactionResponse

	for _, transaction := range transactions {
		transaction := models.TransactionResponse{
			ID:          transaction.ID,
			Total:       transaction.Total,
			Status:      transaction.Status,
			Token:       transaction.Token,
			BookingDate: transaction.BookingDate,
			UserID:      transaction.UserID,
			User:        transaction.User,
			PricingID:   transaction.PricingID,
			Pricing: models.PricingResponse{
				ID:          transaction.Pricing.ID,
				Caption:     transaction.Pricing.Caption,
				Title:       transaction.Pricing.Title,
				Description: transaction.Pricing.Description,
				Image:       transaction.Pricing.Image,
				Content:     transaction.Pricing.Content,
			},
		}
		result = append(result, transaction)
	}
	return result
}

func TimeIn(name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
