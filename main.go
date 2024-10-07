package main

import (
	"fmt"
	"go-login-crud/internal/config"
	"go-login-crud/internal/database"
	"go-login-crud/internal/router"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := config.NewConfigPG()
	dbConn, errdb := database.CreateConnectionPostgres(cfg)
	if errdb != nil {
		fmt.Printf("DB Connection error: %s", errdb)
	}

	r := router.NewRouter(dbConn)
	var addr string = "127.0.0.1:8081"
	if port := os.Getenv("SERVER_PORT"); port != "" {
		addr = ":" + port
	}

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      r,
	}
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}