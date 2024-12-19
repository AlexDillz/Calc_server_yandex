package application

import (
	"log"
	"net/http"
	"os"

	"github.com/AlexDillz/Calc_server_yandex/internal/handlers"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := &Config{
		Addr: os.Getenv("PORT"),
	}
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", handlers.CalcHandler)

	log.Printf("Starting server on port %s...\n", a.config.Addr)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
