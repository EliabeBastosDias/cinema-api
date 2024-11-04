package sessionservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type SessionService struct {
	sessionRepository  ports.SessionRepository
	threaterRepository ports.ThreaterRepository
	movieRepository    ports.MovieRepository
	logger             logger.Provider
}

func New(repo repositories.RepoProvider, logger logger.Provider) *SessionService {
	return &SessionService{
		sessionRepository: repo.Session(),
		logger:            logger,
	}
}

func (s *SessionService) Create(params CreateSessionParams) (*domain.Session, error) {
	return NewCreateSessionCommand(
		s.sessionRepository,
		s.threaterRepository,
		s.movieRepository,
		s.logger,
	).Execute(params)
}

func (s *SessionService) Get(params GetSessionParams) (*domain.Session, error) {
	return NewGetSessionCommand(
		s.sessionRepository,
		s.logger,
	).Execute(params)
}

func (s *SessionService) List(params ListSessionsParams) ([]domain.Session, error) {
	return NewListSessionsCommand(
		s.sessionRepository,
		s.logger,
	).Execute(params)
}

func (s *SessionService) Update(params UpdateSessionParams) error {
	return NewUpdateSessionCommand(
		s.sessionRepository,
		s.logger,
	).Execute(params)
}

func (s *SessionService) Inactivate(params InactivateSessionParams) error {
	return NewInactivateSessionCommand(
		s.sessionRepository,
		s.logger,
	).Execute(params)
}
