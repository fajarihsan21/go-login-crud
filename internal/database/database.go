package database

import (
	"database/sql"
	"go-login-crud/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnectionPostgres(config *config.Config) (*sql.DB, error) {
	conn, errdb := sql.Open("postgres", config.Dsn)
	if errdb != nil {
		log.Fatalf("Error in OpenConnection PGSQL %s", errdb)
		return nil, errdb
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Error in PingConnection PGSQL %s", err)
		return nil, err
	}
	return conn, nil
}