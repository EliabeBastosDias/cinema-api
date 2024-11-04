package threaterservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type UpdateThreaterCommand struct {
	threaterRepository ports.ThreaterRepository
	logger             logger.Provider
}

type UpdateThreaterParams struct {
	ThreaterToken string
	Number        int
	Description   string
}

func NewUpdateThreaterCommand(
	threaterRepository ports.ThreaterRepository,
	logger logger.Provider,
) *UpdateThreaterCommand {
	return &UpdateThreaterCommand{
		threaterRepository: threaterRepository,
		logger:             logger,
	}
}

func (cmd *UpdateThreaterCommand) Execute(params UpdateThreaterParams) error {
	cmd.logger.Info("UpdateThreaterCommand initiated", params)

	threater, err := cmd.threaterRepository.Get(params.ThreaterToken)
	if err != nil {
		cmd.logger.Error("UpdateThreaterCommand failed", err)
		return err
	}

	threater.Number = params.Number
	threater.Description = params.Description

	err = cmd.threaterRepository.Update(threater)
	if err != nil {
		cmd.logger.Error("UpdateThreaterCommand failed", err)
		return err
	}

	cmd.logger.Info("UpdateThreaterCommand finished", params)

	return nil
}
