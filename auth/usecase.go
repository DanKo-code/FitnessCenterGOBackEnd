package auth

import (
	"FitnessCenter_GoBackEnd/dtos"
	"FitnessCenter_GoBackEnd/models"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, signUpDTO dtos.SignUpDTO) (string, string, error)
	SignIn(ctx context.Context, signInDTO dtos.SignInDTO) (*models.User, string, string, error)
	LogOut(refreshToken string) error
}
