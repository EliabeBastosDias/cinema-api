package movieservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type MovieService struct {
	movieRepository   ports.MovieRepository
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

func New(repo repositories.RepoProvider, logger logger.Provider) *MovieService {
	return &MovieService{
		movieRepository:   repo.Movie(),
		sessionRepository: repo.Session(),
		logger:            logger,
	}
}

func (s *MovieService) Create(params CreateMovieParams) (*domain.Movie, error) {
	return NewCreateMovieCommand(
		s.movieRepository,
		s.logger,
	).Execute(params)
}

func (s *MovieService) Get(params GetMovieParams) (*domain.Movie, error) {
	return NewGetMovieCommand(
		s.movieRepository,
		s.logger,
	).Execute(params)
}

func (s *MovieService) List() ([]domain.Movie, error) {
	return NewListMoviesCommand(
		s.movieRepository,
		s.logger,
	).Execute()
}

func (s *MovieService) Update(params UpdateMovieParams) error {
	return NewUpdateMovieCommand(
		s.movieRepository,
		s.logger,
	).Execute(params)
}

func (s *MovieService) Inactivate(params InactivateMovieParams) error {
	return NewInactivateMovieCommand(
		s.movieRepository,
		s.sessionRepository,
		s.logger,
	).Execute(params)
}
