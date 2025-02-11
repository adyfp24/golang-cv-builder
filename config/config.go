package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error load .enc file")
	}
}

func GetEnv(key string) string{
	return os.Getenv(key)
}