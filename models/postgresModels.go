package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	FirstName   string    `gorm:"column:first_name"`
	LastName    string    `gorm:"column:last_name"`
	Email       string    `gorm:"column:email"`
	Role        string    `gorm:"column:role"`
	Password    string    `gorm:"column:password_hash"`
	Photo       string    `gorm:"column:photo"`
	CreatedTime time.Time `gorm:"column:created_time; autoCreateTime"`
}
