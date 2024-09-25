package main

import (
	authpostgres "FitnessCenter_GoBackEnd/auth/repository/postgres"
	"FitnessCenter_GoBackEnd/config"
	"FitnessCenter_GoBackEnd/server"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	db := server.InitDB()
	pr := authpostgres.NewUserRepository(db)
	client, err := pr.GetUserByEmail("jane.smith@example.com")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("Client: ", client)
}
