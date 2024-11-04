package sessionservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type ListSessionsCommand struct {
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

type ListSessionsParams struct {
	Page int
}

func NewListSessionsCommand(SessionRepo ports.SessionRepository, logger logger.Provider) *ListSessionsCommand {
	return &ListSessionsCommand{
		sessionRepository: SessionRepo,
		logger:            logger,
	}
}

func (cmd *ListSessionsCommand) Execute(params ListSessionsParams) ([]domain.Session, error) {
	cmd.logger.Info("ListSessionsCommand initiated")

	sessions, err := cmd.sessionRepository.List(params.Page)
	if err != nil {
		cmd.logger.Error("ListSessionsCommand initiated", err)
		return nil, fmt.Errorf("error creating Session: %w", err)
	}

	cmd.logger.Info("ListSessionsCommand finished")

	return sessions, nil
}
