package server

import (
	"FitnessCenter_GoBackEnd/auth"
	authhttp "FitnessCenter_GoBackEnd/auth/delivery/http"
	authpostgres "FitnessCenter_GoBackEnd/auth/repository/postgres"
	authusecase "FitnessCenter_GoBackEnd/auth/usecase"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authpostgres.NewUserRepository(db)

	return &App{
		authUC: authusecase.NewAuthUseCase(userRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *gorm.DB {
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
