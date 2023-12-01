package handlers

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"
	dto "wedding/dto"
	"wedding/models"
	"wedding/pkg/bcrypt"
	jwtToken "wedding/pkg/jwt"
	"wedding/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var path_photo_auth = "http://localhost:5000/uploads/photo/"

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

// function random id user
func generateRandomID() int {
	rand.Seed(time.Now().UnixNano())
	min := 10000000000 // Min 11 digit
	max := 99999999999 // Max 11 digit
	return rand.Intn(max-min+1) + min
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(dto.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	randomID := generateRandomID()

	checkUser, err := h.AuthRepository.FindUserByUsernameOrEmail(request.Username, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	if checkUser.ID != 0 {
		errorMessage := "Username or email already exists."
		return c.JSON(http.StatusConflict, dto.ErrorResult{Status: http.StatusConflict, Message: errorMessage})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		ID:       randomID,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Role:     "user",
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerAuth) RegisterAdmin(c echo.Context) error {
	request := new(dto.RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	hashedPassword, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(dto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "wrong email or password"})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	user.Photo = path_photo_auth + user.Photo

	loginResponse := dto.LoginResponse{
		// ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	user.Photo = path_photo_auth + user.Photo

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ConvertAuthResponse(user)})
}

func ConvertAuthResponse(user models.User) models.UserResponse {
	return models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Address:  user.Address,
		Photo:    user.Photo,
	}
}
