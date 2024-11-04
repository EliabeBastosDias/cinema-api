package sessionservice

import (
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type CreateSessionCommand struct {
	sessionRepository  ports.SessionRepository
	threaterRepository ports.ThreaterRepository
	movieRepository    ports.MovieRepository
	logger             logger.Provider
}

type CreateSessionParams struct {
	MovieToken      string
	ThreaterToken   string
	SessionDatetime time.Time
}

func NewCreateSessionCommand(sessionRepo ports.SessionRepository, threaterRepo ports.ThreaterRepository, movieRepo ports.MovieRepository, logger logger.Provider) *CreateSessionCommand {
	return &CreateSessionCommand{
		sessionRepository:  sessionRepo,
		threaterRepository: threaterRepo,
		movieRepository:    movieRepo,
		logger:             logger,
	}
}

func (cmd *CreateSessionCommand) Execute(params CreateSessionParams) (*domain.Session, error) {
	cmd.logger.Info("CreateSessionCommand initiated", params)

	if _, err := cmd.threaterRepository.Get(params.ThreaterToken); err != nil {
		cmd.logger.Error("CreateSessionCommand failed", err)
		return nil, err
	}

	if _, err := cmd.movieRepository.Get(params.ThreaterToken); err != nil {
		cmd.logger.Error("CreateSessionCommand failed", err)
		return nil, err
	}

	sessionToCreate := domain.Session{
		MovieToken:      params.MovieToken,
		ThreaterToken:   params.ThreaterToken,
		SessionDatetime: params.SessionDatetime,
	}

	createdSession, err := cmd.sessionRepository.Create(&sessionToCreate)
	if err != nil {
		cmd.logger.Info("CreateSessionCommand failed", err)
		return nil, err
	}

	cmd.logger.Info("CreateSessionCommand finished", params)

	return createdSession, nil
}
