package http

import (
	"FitnessCenter_GoBackEnd/auth"
	"FitnessCenter_GoBackEnd/constants"
	"FitnessCenter_GoBackEnd/dtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	useCase  auth.UseCase
	validate *validator.Validate
}

func NewHandler(useCase auth.UseCase, validator *validator.Validate) *Handler {
	return &Handler{
		useCase:  useCase,
		validate: validator,
	}
}

func (h *Handler) SignUp(c *gin.Context) {

	var inp dtos.SignUpDTO

	if err := c.ShouldBindJSON(&inp); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fingerPrintValue, exists := c.Get("fingerprint")
	if !exists {
		log.Println("fingerPrint not found in context")
		return
	}
	FingerPrintValueCasted, ok := fingerPrintValue.(string)
	if !ok {
		log.Println("fingerPrint is not a valid string")
		return
	}
	inp.FingerPrint = FingerPrintValueCasted

	err := h.validate.Struct(inp)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(), // Optional: Include the actual validation error message
		})
		return
	}

	accessToken, refreshToken, err := h.useCase.SignUp(c.Request.Context(), inp)
	if err != nil {
		log.Println(err)
		return
	}

	c.SetCookie(
		"refreshToken",
		refreshToken,
		constants.COOKIE_SETTINGS.RefreshToken.MaxAge,
		"",
		"",
		false,
		constants.COOKIE_SETTINGS.RefreshToken.HttpOnly,
	)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":           accessToken,
		"accessTokenExpiration": constants.ACCESS_TOKEN_EXPIRATION,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var inp dtos.SignInDTO

	if err := c.ShouldBindJSON(&inp); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fingerPrintValue, exists := c.Get("fingerprint")
	if !exists {
		log.Println("fingerPrint not found in context")
		return
	}
	FingerPrintValueCasted, ok := fingerPrintValue.(string)
	if !ok {
		log.Println("fingerPrint is not a valid string")
		return
	}
	inp.FingerPrint = FingerPrintValueCasted

	err := h.validate.Struct(inp)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(), // Optional: Include the actual validation error message
		})
		return
	}

	user, accessToken, refreshToken, err := h.useCase.SignIn(c.Request.Context(), inp)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return a 404 Not Found if the user doesn't exist
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}

		// For other errors (e.g., database issues), return a 500 Internal Server Error
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.SetCookie(
		"refreshToken",
		refreshToken,
		constants.COOKIE_SETTINGS.RefreshToken.MaxAge,
		"",
		"",
		false,
		constants.COOKIE_SETTINGS.RefreshToken.HttpOnly,
	)

	c.JSON(http.StatusOK, gin.H{
		"user":                  user,
		"accessToken":           accessToken,
		"accessTokenExpiration": constants.ACCESS_TOKEN_EXPIRATION,
	})
}

func (h *Handler) LogOut(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.LogOut(refreshToken); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"err: ": err})
	}

	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
}
