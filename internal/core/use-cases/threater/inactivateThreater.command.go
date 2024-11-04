package threaterservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type InactivateThreaterCommand struct {
	threaterRepository ports.ThreaterRepository
	sessionRepository  ports.SessionRepository
	logger             logger.Provider
}

type InactivateThreaterParams struct {
	ThreaterToken string
}

func NewInactivateThreaterCommand(
	threaterRepository ports.ThreaterRepository,
	sessionRepository ports.SessionRepository,
	logger logger.Provider,
) *InactivateThreaterCommand {
	return &InactivateThreaterCommand{
		sessionRepository:  sessionRepository,
		threaterRepository: threaterRepository,
		logger:             logger,
	}
}

func (cmd *InactivateThreaterCommand) Execute(params InactivateThreaterParams) error {
	cmd.logger.Info("InactivateThreaterCommand initiated", params)

	_, err := cmd.threaterRepository.Get(params.ThreaterToken)
	if err != nil {
		cmd.logger.Error("InactivateThreaterCommand failed", err)
		return err
	}

	err = cmd.threaterRepository.Inactivate(params.ThreaterToken)
	if err != nil {
		cmd.logger.Error("InactivateThreaterCommand failed", err)
		return err
	}

	cmd.logger.Info("InactivateThreaterCommand finished", params)

	return nil
}
