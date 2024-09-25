package models

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID        uuid.UUID `gorm:"primarykey"`
	FirstName string    `gorm:"column:firstname"`
	LastName  string    `gorm:"column:lastname"`
	Email     string    `gorm:"column:email"`
	Role      int       `gorm:"column:role"`
	Password  string    `gorm:"column:password"`
	Photo     string    `gorm:"column:photo"`
	CreatedAt time.Time `gorm:"column:createdat"`
}
