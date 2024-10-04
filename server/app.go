package server

import (
	"FitnessCenter_GoBackEnd/auth"
	authhttp "FitnessCenter_GoBackEnd/auth/delivery/http"
	authpostgres "FitnessCenter_GoBackEnd/auth/repository/postgres"
	authusecase "FitnessCenter_GoBackEnd/auth/usecase"
	"FitnessCenter_GoBackEnd/validators"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	refreshSessionRepo := authpostgres.NewRefreshSessionRepository(db)

	return &App{
		authUC: authusecase.NewAuthUseCase(userRepo, refreshSessionRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                            // Allow requests from this origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow cookies (if needed)
		MaxAge:           12 * time.Hour, // How long the results of a preflight request can be cached
	}))

	validate := validator.New()
	err := validate.RegisterValidation("validFirstName", validators.ValidateUsreFisrtName)
	if err != nil {
		return nil
	}

	authhttp.RegisterHTTPEndpoints(router, a.authUC, validate)

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
}
