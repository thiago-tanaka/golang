package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionDB = "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"
	Port         = 0
)

func LoadConfiguration() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	ConnectionDB = os.Getenv("MYSQL_CONNECTION")
}
