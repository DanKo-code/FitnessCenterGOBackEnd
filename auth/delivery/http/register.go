package http

import (
	"FitnessCenter_GoBackEnd/auth"
	fingerprintMiddleware "FitnessCenter_GoBackEnd/auth/delivery/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase, validator *validator.Validate) {
	h := NewHandler(uc, validator)

	authEndPoints := router.Group("/auth")
	{
		authEndPoints.Use(fingerprintMiddleware.FingerprintMiddleware())
		authEndPoints.POST("/signUp", h.SignUp)
		authEndPoints.POST("/signIn", h.SignIn)
		authEndPoints.POST("/logOut", h.LogOut)
	}

}
