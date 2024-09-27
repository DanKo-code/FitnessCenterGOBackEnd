package usecase

import (
	"FitnessCenter_GoBackEnd/auth"
	"FitnessCenter_GoBackEnd/models"
	"context"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

func NewAuthUseCase(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (a *AuthUseCase) SignUp(ctx context.Context, FirstName string, LastName string, Email string, password string) error {

	user := new(models.User)

	user, err := a.userRepo.GetUserByEmail(Email)
	if err == nil {
		return auth.ErrUserAlreadyExists
	}

	_ = user

	return nil
}
