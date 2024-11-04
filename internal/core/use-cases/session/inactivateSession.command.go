package sessionservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type InactivateSessionCommand struct {
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

type InactivateSessionParams struct {
	SessionToken string
}

func NewInactivateSessionCommand(
	sessionRepository ports.SessionRepository,
	logger logger.Provider,
) *InactivateSessionCommand {
	return &InactivateSessionCommand{
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}

func (cmd *InactivateSessionCommand) Execute(params InactivateSessionParams) error {
	cmd.logger.Info("InactivateSessionCommand initiated", params)

	_, err := cmd.sessionRepository.Get(params.SessionToken)
	if err != nil {
		cmd.logger.Error("InactivateSessionCommand failed", err)
		return err
	}

	err = cmd.sessionRepository.Inactivate(params.SessionToken)
	if err != nil {
		cmd.logger.Error("InactivateSessionCommand failed", err)
		return err
	}

	cmd.logger.Info("InactivateSessionCommand finished", params)

	return nil
}
