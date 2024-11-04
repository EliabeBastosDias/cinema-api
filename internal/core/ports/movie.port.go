package ports

import "github.com/EliabeBastosDias/cinema-api/internal/core/domain"

type MovieRepository interface {
	Create(movie *domain.Movie) (*domain.Movie, error)
	Get(movieToken string) (*domain.Movie, error)
	List() ([]domain.Movie, error)
	Update(movieToUpdate *domain.Movie) error
	Inactivate(movieToken string) error
}

type MovieService interface{}
