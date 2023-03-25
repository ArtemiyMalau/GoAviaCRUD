package main

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type config struct {
	ListenAddr string `env:"LISTEN_ADDR" envDefault:":3000"`

	DB struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Database string `env:"DB_DATABASE"`
		Username string `env:"DB_USERNAME"`
		Password string `env:"DB_PASSWORD"`

		Scripts string `env:"DB_SCRIPTS_PATH"`
	}
}

var once sync.Once
var configInstance *config

func GetConfig() *config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("unable to load .env file: %e", err)
		}

		configInstance = &config{}
		if err := env.Parse(configInstance); err != nil {
			log.Fatalf("unable to parse .env file: %e", err)
		}
	})
	return configInstance
}
