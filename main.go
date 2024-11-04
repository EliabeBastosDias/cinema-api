package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/EliabeBastosDias/cinema-api/cmd/http"
	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	"github.com/EliabeBastosDias/cinema-api/internal/config"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/authhdl"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/docs"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/healthy"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/moviehdl"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/sessionhdl"
	"github.com/EliabeBastosDias/cinema-api/internal/handlers/threaterhdl"
)

func main() {
	conf := config.New()
	apt := adapters.New(conf)

	app := server.New(apt)

	app.RegisterHandler(
		docs.NewHandler(apt),
		authhdl.NewHandler(apt),
		healthy.NewHandler(apt),
		moviehdl.NewHandler(apt),
		sessionhdl.NewHandler(apt),
		threaterhdl.NewHandler(apt),
	)

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := app.Run()
		if err != nil {
			stopCh <- syscall.SIGTERM
		}
	}()
	<-stopCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	apt.Shutdown(ctx)
	app.Shutdown(ctx)

	err := app.Router.Run(app.Server.Addr)
	if err != nil {
		panic(fmt.Sprintf("Unable to listen server: %v", err))
	}
}
