package threaterservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type ThreaterService struct {
	threaterRepository ports.ThreaterRepository
	sessionRepository  ports.SessionRepository
	logger             logger.Provider
}

func New(repo repositories.RepoProvider, logger logger.Provider) *ThreaterService {
	return &ThreaterService{
		threaterRepository: repo.Threater(),
		sessionRepository:  repo.Session(),
		logger:             logger,
	}
}

func (s *ThreaterService) Create(params CreateThreaterParams) (*domain.Threater, error) {
	return NewCreateThreaterCommand(
		s.threaterRepository,
		s.logger,
	).Execute(params)
}

func (s *ThreaterService) Get(params GetThreaterParams) (*domain.Threater, error) {
	return NewGetThreaterCommand(
		s.threaterRepository,
		s.logger,
	).Execute(params)
}

func (s *ThreaterService) List() ([]domain.Threater, error) {
	return NewListThreatersCommand(
		s.threaterRepository,
		s.logger,
	).Execute()
}

func (s *ThreaterService) Update(params UpdateThreaterParams) error {
	return NewUpdateThreaterCommand(
		s.threaterRepository,
		s.logger,
	).Execute(params)
}

func (s *ThreaterService) Inactivate(params InactivateThreaterParams) error {
	return NewInactivateThreaterCommand(
		s.threaterRepository,
		s.sessionRepository,
		s.logger,
	).Execute(params)
}
