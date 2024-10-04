package models

import (
	"github.com/google/uuid"
	"time"
)

type Abonement struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title        string    `gorm:"type:varchar(255)"`
	Validity     string    `gorm:"type:varchar(255)"`
	VisitingTime string    `gorm:"type:varchar(255)"`
	Photo        string    `gorm:"type:varchar(255)"`
	Price        int
	Services     []*Service `gorm:"many2many:abonement_service;"`
}

type Service struct {
	ID         uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Title      string       `gorm:"type:varchar(255)"`
	Photo      string       `gorm:"type:varchar(255)"`
	Abonements []*Abonement `gorm:"many2many:abonement_service;"`
	Coaches    []*Coach     `gorm:"many2many:coach_service;"`
}

type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName       string    `gorm:"type:varchar(255);not null"`
	LastName        string    `gorm:"type:varchar(255);not null"`
	Email           string    `gorm:"type:varchar(255);not null"`
	Role            string    `gorm:"type:varchar(255);not null"`
	PasswordHash    string    `gorm:"type:varchar(255);not null"`
	Photo           string    `gorm:"type:varchar(255)"`
	CreatedTime     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Comments        []Comment
	Orders          []Order
	RefreshSessions []RefreshSession
}

type Coach struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255);not null"`
	Photo       string     `gorm:"type:varchar(255)"`
	Services    []*Service `gorm:"many2many:coach_service;"`
	Comments    []Comment
}

type Comment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CommentBody string    `gorm:"type:varchar(255);not null"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	CoachID     uuid.UUID `gorm:"type:uuid;not null"`
	CreateDate  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Coach       Coach     `gorm:"foreignKey:CoachID;constraint:OnDelete:CASCADE;"`
}

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	AbonementID uuid.UUID `gorm:"type:uuid"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	Status      int       `gorm:"not null"`
	Abonement   Abonement `gorm:"foreignKey:AbonementID"`
	User        User      `gorm:"foreignKey:UserID"`
}

type RefreshSession struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	RefreshToken string    `gorm:"type:varchar(400);not null"`
	FingerPrint  string    `gorm:"type:varchar(32);not null"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
