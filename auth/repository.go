package auth

import (
	"FitnessCenter_GoBackEnd/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}
