package server

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	httpServer *http.Server

	//authUC _
}

func NewApp() *App {
	return &App{}
}

func InitDB() {
	dsn := "host=localhost user=FitnessCenterUser password=tanki@danik2003 dbname=fitnesscenterdb port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
