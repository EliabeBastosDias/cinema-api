package sessionservice

import (
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type UpdateSessionCommand struct {
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

type UpdateSessionParams struct {
	SessionToken    string
	MovieToken      string
	ThreaterToken   string
	SessionDatetime time.Time
}

func NewUpdateSessionCommand(
	SessionRepository ports.SessionRepository,
	logger logger.Provider,
) *UpdateSessionCommand {
	return &UpdateSessionCommand{
		sessionRepository: SessionRepository,
		logger:            logger,
	}
}

func (cmd *UpdateSessionCommand) Execute(params UpdateSessionParams) error {
	cmd.logger.Info("ListSessionsCommand initiated", params)

	session, err := cmd.sessionRepository.Get(params.SessionToken)
	if err != nil {
		cmd.logger.Error("ListSessionsCommand failed", err)
		return err
	}

	session.MovieToken = params.MovieToken
	session.ThreaterToken = params.ThreaterToken
	session.SessionDatetime = params.SessionDatetime

	err = cmd.sessionRepository.Update(session)
	if err != nil {
		cmd.logger.Error("ListSessionsCommand failed", err)
		return err
	}

	return nil
}
