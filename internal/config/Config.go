package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn string
}

func NewConfigPG() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	config := new(Config)
	server := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PSWD")
	dbname := os.Getenv("PG_NAME")
	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server,
		port,
		user,
		pass,
		dbname)
	config.Dsn = postgresInfo
	return config
}

// func CreateConnectionPg() *Config {
// 	return CreateConnectionPostgres(os.Getenv("PG_NAME"))
// }