package main

import (
	"FitnessCenter_GoBackEnd/config"
	"FitnessCenter_GoBackEnd/server"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
	/*db := server.InitDB()

	pr := authpostgres.NewUserRepository(db)

	client, err := pr.GetUserByEmail("jane.smith@example.com")

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Printf("Client: %+v", client)
	*/
}
