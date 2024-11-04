package sessionservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type GetSessionCommand struct {
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

type GetSessionParams struct {
	SessionToken string
}

func NewGetSessionCommand(sessionRepo ports.SessionRepository, logger logger.Provider) *GetSessionCommand {
	return &GetSessionCommand{
		sessionRepository: sessionRepo,
		logger:            logger,
	}
}

func (cmd *GetSessionCommand) Execute(params GetSessionParams) (*domain.Session, error) {
	cmd.logger.Info("GetSessionCommand initiated", params)

	session, err := cmd.sessionRepository.Get(params.SessionToken)
	if err != nil {
		cmd.logger.Error("GetSessionCommand failed", err)
		return nil, fmt.Errorf("error creating session: %w", err)
	}

	cmd.logger.Info("GetSessionCommand finished", params)

	return session, nil
}
