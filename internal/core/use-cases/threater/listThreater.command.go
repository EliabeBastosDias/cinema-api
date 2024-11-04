package threaterservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type ListThreaterCommand struct {
	threaterRepository ports.ThreaterRepository
	logger             logger.Provider
}

func NewListThreatersCommand(threaterRepo ports.ThreaterRepository, logger logger.Provider) *ListThreaterCommand {
	return &ListThreaterCommand{
		threaterRepository: threaterRepo,
		logger:             logger,
	}
}

func (cmd *ListThreaterCommand) Execute() ([]domain.Threater, error) {
	cmd.logger.Info("ListThreaterCommand initiated")

	threaters, err := cmd.threaterRepository.List()
	if err != nil {
		cmd.logger.Error("ListThreaterCommand failed", err)
		return nil, fmt.Errorf("error creating Threater: %w", err)
	}

	cmd.logger.Info("ListThreaterCommand finished")

	return threaters, nil
}
