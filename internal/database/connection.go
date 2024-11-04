package database

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

func Connect(dsn string) *dbr.Session {
	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		panic(fmt.Sprintf("Could not open database connection: %v", err))
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err := conn.Ping(); err != nil {
		panic(fmt.Sprintf("Unable to ping database: %v", err))
	}

	fmt.Println("🐘 Postgres connection established.")

	return conn.NewSession(nil)
}
