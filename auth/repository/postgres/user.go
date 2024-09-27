package postgres

import (
	"FitnessCenter_GoBackEnd/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var client models.User
	if err := ur.db.Where("Email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (ur *UserRepository) CreateUser(user models.User) (*models.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
