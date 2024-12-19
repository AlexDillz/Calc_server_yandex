package main

import (
	"log"

	"github.com/AlexDillz/Calc_server_yandex/internal/application"
)

func main() {

	app := application.New()

	if err := app.RunServer(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
