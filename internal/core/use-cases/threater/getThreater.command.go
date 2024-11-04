package threaterservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type GetThreaterCommand struct {
	threaterRepository ports.ThreaterRepository
	logger             logger.Provider
}

type GetThreaterParams struct {
	ThreaterToken string
}

func NewGetThreaterCommand(threaterRepo ports.ThreaterRepository, logger logger.Provider) *GetThreaterCommand {
	return &GetThreaterCommand{
		threaterRepository: threaterRepo,
		logger:             logger,
	}
}

func (cmd *GetThreaterCommand) Execute(params GetThreaterParams) (*domain.Threater, error) {
	cmd.logger.Info("GetThreaterCommand initiated", params)

	createdThreater, err := cmd.threaterRepository.Get(params.ThreaterToken)
	if err != nil {
		cmd.logger.Error("GetThreaterCommand failed", err)
		return nil, fmt.Errorf("error creating threater: %w", err)
	}

	cmd.logger.Info("GetThreaterCommand finished", params)

	return createdThreater, nil
}
