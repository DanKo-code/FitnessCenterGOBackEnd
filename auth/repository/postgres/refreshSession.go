package postgres

import (
	"FitnessCenter_GoBackEnd/models"
	"gorm.io/gorm"
)

type RefreshSessionRepository struct {
	db *gorm.DB
}

func NewRefreshSessionRepository(db *gorm.DB) *RefreshSessionRepository {
	return &RefreshSessionRepository{db: db}
}

func (rsr *RefreshSessionRepository) CreateRefreshSession(refreshSession models.RefreshSession) (*models.RefreshSession, error) {
	if err := rsr.db.Create(&refreshSession).Error; err != nil {
		return nil, err
	}

	return &refreshSession, nil
}

func (rsr *RefreshSessionRepository) DeleteRefreshSession(refreshToken string) error {
	if err := rsr.db.Delete(&models.RefreshSession{}, "refresh_token = ?", refreshToken).Error; err != nil {
		return err
	}

	return nil
}

func (rsr *RefreshSessionRepository) GetRefreshSession(refreshToken string) (*models.RefreshSession, error) {
	refreshSession := &models.RefreshSession{}
	if err := rsr.db.Where("refresh_token = ?", refreshToken).Find(refreshSession).Error; err != nil {
		return nil, err
	}

	return refreshSession, nil
}
