package adapters

import (
	"context"
	"io"
	"log"

	"github.com/EliabeBastosDias/cinema-api/internal/config"
	"github.com/EliabeBastosDias/cinema-api/internal/database"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gocraft/dbr/v2"
	"github.com/opensearch-project/opensearch-go"
	"github.com/streadway/amqp"
)

func New(conf config.Config) *Adapters {
	a := &Adapters{}
	a.Config = conf

	l := logger.New()
	a.Logger = l

	session := database.Connect(conf.DSN)
	a.DB = session
	a.closers = append(a.closers, session)

	rabbitConn, err := amqp.Dial(conf.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	a.RabbitMQConn = rabbitConn
	a.closers = append(a.closers, rabbitConn)

	opensearchClient, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{conf.OpenSearchURL},
	})
	if err != nil {
		log.Fatalf("Failed to connect to OpenSearch: %v", err)
	}
	a.OpenSearch = opensearchClient

	a.Repositories = repositories.New(a.DB)

	a.closers = append(a.closers, l)

	return a
}

type Adapters struct {
	DB           *dbr.Session
	Config       config.Config
	Logger       logger.Provider
	closers      []io.Closer
	Repositories repositories.RepoProvider

	RabbitMQConn *amqp.Connection
	OpenSearch   *opensearch.Client
}

func (a *Adapters) Shutdown(ctx context.Context) {
	for _, c := range a.closers {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.Close()
			if err != nil {
				log.Printf("Error closing component: %v", err)
			}
		}
	}
}
