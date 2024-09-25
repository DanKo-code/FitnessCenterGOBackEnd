package server

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

type App struct {
	httpServer *http.Server

	//authUC _
}

func NewApp() *App {
	return &App{}
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString("postgresDNS.host"),
		viper.GetString("postgresDNS.user"),
		viper.GetString("postgresDNS.password"),
		viper.GetString("postgresDNS.dbname"),
		viper.GetString("postgresDNS.port"),
		viper.GetString("postgresDNS.sslmode"),
		viper.GetString("postgresDNS.TimeZone"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
	/*result, err := db.First(&Client{}).Rows()
	if err != nil {
		fmt.Println("Error: ", result)
	}

	for result.Next() {
		var client Client
		db.ScanRows(result, &client) // Scan each row into a client
		fmt.Printf("Client: %+v\n", client)
	}*/
}
