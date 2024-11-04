package config

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/envs"
)

type Config struct {
	DSN           string
	Port          string
	RabbitMQURL   string
	OpenSearchURL string
}

func New() Config {
	envs.Load()
	port := envs.Get(envs.SERVER_PORT)

	dbUsername := envs.Get(envs.DB_USERNAME)
	dbPassword := envs.Get(envs.DB_PASSWORD)
	dbHost := envs.Get(envs.DB_HOST)
	dbPort := envs.Get(envs.DB_PORT)
	dbName := envs.Get(envs.DB_NAME)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)

	rabbitMQURL := envs.Get(envs.RABBITMQ_URL)
	openSearchURL := envs.Get(envs.OPENSEARCH_URL)

	return Config{
		DSN:           dsn,
		Port:          port,
		RabbitMQURL:   rabbitMQURL,
		OpenSearchURL: openSearchURL,
	}
}
