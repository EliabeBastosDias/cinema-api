package movieservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type ListMovieCommand struct {
	movieRepository ports.MovieRepository
	logger          logger.Provider
}

func NewListMoviesCommand(MovieRepo ports.MovieRepository, logger logger.Provider) *ListMovieCommand {
	return &ListMovieCommand{
		movieRepository: MovieRepo,
		logger:          logger,
	}
}

func (cmd *ListMovieCommand) Execute() ([]domain.Movie, error) {
	cmd.logger.Info("ListMovieCommand initiated")

	movies, err := cmd.movieRepository.List()
	if err != nil {
		cmd.logger.Error("ListMovieCommand failed", err)
		return nil, fmt.Errorf("error creating movie: %w", err)
	}

	cmd.logger.Info("ListMovieCommand finished")

	return movies, nil
}
