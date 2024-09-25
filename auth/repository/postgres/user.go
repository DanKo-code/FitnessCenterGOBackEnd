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

func (ur *UserRepository) GetUserByEmail(email string) (*models.Client, error) {
	var client models.Client
	result, err := ur.db.Where("Email = ?", email).First(&models.Client{}).Rows()
	if err != nil {
		return nil, err
	}

	for result.Next() {
		if err := ur.db.ScanRows(result, &client); err != nil {
			return nil, err
		}
	}

	return &client, nil
}
