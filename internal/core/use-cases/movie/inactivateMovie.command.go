package movieservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type InactivateMovieCommand struct {
	movieRepository   ports.MovieRepository
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

type InactivateMovieParams struct {
	MovieToken string
}

func NewInactivateMovieCommand(
	movieRepository ports.MovieRepository,
	sessionRepository ports.SessionRepository,
	logger logger.Provider,
) *InactivateMovieCommand {
	return &InactivateMovieCommand{
		movieRepository:   movieRepository,
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}

func (cmd *InactivateMovieCommand) Execute(params InactivateMovieParams) error {
	cmd.logger.Info("InactivateMovieCommand initiated", params)

	_, err := cmd.movieRepository.Get(params.MovieToken)
	if err != nil {
		cmd.logger.Error("InactivateMovieCommand failed", err)
		return err
	}

	err = cmd.movieRepository.Inactivate(params.MovieToken)
	if err != nil {
		cmd.logger.Error("InactivateMovieCommand failed", err)
		return err
	}

	cmd.logger.Info("InactivateMovieCommand finished", params)

	return nil
}
