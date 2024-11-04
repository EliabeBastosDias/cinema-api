package movieservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type CreateMovieCommand struct {
	movieRepository ports.MovieRepository
	logger          logger.Provider
}

type CreateMovieParams struct {
	Name     string
	Director string
	Duration int
}

func NewCreateMovieCommand(movieRepo ports.MovieRepository, logger logger.Provider) *CreateMovieCommand {
	return &CreateMovieCommand{
		movieRepository: movieRepo,
		logger:          logger,
	}
}

func (cmd *CreateMovieCommand) Execute(params CreateMovieParams) (*domain.Movie, error) {
	cmd.logger.Info("CreateMovieCommand initiated", params)

	createdMovie, err := cmd.movieRepository.Create(&domain.Movie{
		Name:     params.Name,
		Director: params.Director,
		Duration: params.Duration,
	})
	if err != nil {
		cmd.logger.Error("CreateMovieCommand failed", err)
		return nil, fmt.Errorf("error creating movie: %w", err)
	}

	cmd.logger.Info("CreateMovieCommand finished", params)

	return createdMovie, nil
}
