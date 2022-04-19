package config

import (
	"github.com/joho/godotenv"
)

func init() {
	loadConfig()
}

func loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
