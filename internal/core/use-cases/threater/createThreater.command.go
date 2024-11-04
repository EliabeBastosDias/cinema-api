package threaterservice

import (
	"fmt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type CreateThreaterCommand struct {
	threaterRepository ports.ThreaterRepository
	logger             logger.Provider
}

type CreateThreaterParams struct {
	Number      int
	Description string
}

func NewCreateThreaterCommand(threaterRepo ports.ThreaterRepository, logger logger.Provider) *CreateThreaterCommand {
	return &CreateThreaterCommand{
		threaterRepository: threaterRepo,
		logger:             logger,
	}
}

func (cmd *CreateThreaterCommand) Execute(params CreateThreaterParams) (*domain.Threater, error) {
	cmd.logger.Info("CreateThreaterCommand initiated", params)

	threaterToCreate := domain.Threater{
		Number:      params.Number,
		Description: params.Description,
	}

	createdThreater, err := cmd.threaterRepository.Create(&threaterToCreate)
	if err != nil {
		cmd.logger.Error("CreateThreaterCommand failed", err)
		return nil, fmt.Errorf("error creating movie: %w", err)
	}

	cmd.logger.Info("CreateThreaterCommand failed", params)

	return createdThreater, nil
}
