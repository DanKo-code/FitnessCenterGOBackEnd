package auth

import (
	"FitnessCenter_GoBackEnd/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
}

type RefreshSessionRepository interface {
	CreateRefreshSession(refreshSession models.RefreshSession) (*models.RefreshSession, error)
	DeleteRefreshSession(refreshToken string) error
}
