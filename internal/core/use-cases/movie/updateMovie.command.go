package movieservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type UpdateMovieCommand struct {
	movieRepository ports.MovieRepository
	logger          logger.Provider
}

type UpdateMovieParams struct {
	MovieToken string
	Name       string
	Director   string
	Duration   int
}

func NewUpdateMovieCommand(
	movieRepository ports.MovieRepository,
	logger logger.Provider,
) *UpdateMovieCommand {
	return &UpdateMovieCommand{
		movieRepository: movieRepository,
		logger:          logger,
	}
}

func (cmd *UpdateMovieCommand) Execute(params UpdateMovieParams) error {
	cmd.logger.Info("UpdateMovieCommand initiated", params)

	movie, err := cmd.movieRepository.Get(params.MovieToken)
	if err != nil {
		cmd.logger.Error("UpdateMovieCommand failed", err)
		return err
	}

	movie.Name = params.Name
	movie.Director = params.Director
	movie.Duration = params.Duration

	err = cmd.movieRepository.Update(movie)
	if err != nil {
		cmd.logger.Error("UpdateMovieCommand failed", err)
		return err
	}

	cmd.logger.Info("UpdateMovieCommand finished", params)

	return nil
}
