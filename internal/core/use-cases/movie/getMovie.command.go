package movieservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type GetMovieCommand struct {
	movieRepository ports.MovieRepository
	logger          logger.Provider
}

type GetMovieParams struct {
	MovieToken string
}

func NewGetMovieCommand(movieRepo ports.MovieRepository, logger logger.Provider) *GetMovieCommand {
	return &GetMovieCommand{
		movieRepository: movieRepo,
		logger:          logger,
	}
}

func (cmd *GetMovieCommand) Execute(params GetMovieParams) (*domain.Movie, error) {
	cmd.logger.Info("GetMovieCommand initiated", params)

	movie, err := cmd.movieRepository.Get(params.MovieToken)
	if err != nil {
		cmd.logger.Error("GetMovieCommand failed", err)
		return nil, fmt.Errorf("error creating movie: %w", err)

	}

	cmd.logger.Info("GetMovieCommand finished", params)

	return movie, nil
}
