package main

import (
	"FitnessCenter_GoBackEnd/config"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
