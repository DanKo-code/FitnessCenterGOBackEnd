package http

import (
	"FitnessCenter_GoBackEnd/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndPoints := router.Group("/auth")
	{
		authEndPoints.POST("/sign-up", h.SignUp)
	}
}
