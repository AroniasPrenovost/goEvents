package config 

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Env struct {
	Port string // int? 
	DB_port string // int? 
	DB_name string
	DB_user string
	DB_password string
	Admin_password string
}

func InitEnv() (Env) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ENV := Env{}
	ENV.Port = os.Getenv("PORT")
	ENV.DB_port = os.Getenv("DB_PORT")
	ENV.DB_user = os.Getenv("DB_USER")
	ENV.DB_name = os.Getenv("DB_NAME")
	ENV.DB_password = os.Getenv("DB_PASSWORD")
	ENV.Admin_password = os.Getenv("ADMIN_PASSWORD")

	return ENV
}