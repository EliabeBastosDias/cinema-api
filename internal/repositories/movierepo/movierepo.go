package movierepo

import (
	"sync"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type MovieRepository struct {
	baseRepo *baserepo.BaseRepository[domain.Movie]
}

var (
	instance *MovieRepository
	once     sync.Once
)

func New(db *dbr.Session) *MovieRepository {
	once.Do(func() {
		instance = &MovieRepository{
			baseRepo: baserepo.New[domain.Movie](db),
		}
	})
	return instance
}

func (r *MovieRepository) Create(movie *domain.Movie) (*domain.Movie, error) {
	movie.Token = uuid.New()
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	err := r.baseRepo.Insert("movies", movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepository) List() ([]domain.Movie, error) {
	movies, err := r.baseRepo.List("movies")
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) Get(movieToken string) (*domain.Movie, error) {
	movie, err := r.baseRepo.Get("movies", "token", movieToken)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepository) Update(movieToUpdate *domain.Movie) error {
	updates := map[string]interface{}{
		"name":       movieToUpdate.Name,
		"director":   movieToUpdate.Director,
		"duration":   movieToUpdate.Duration,
		"updated_at": time.Now(),
	}

	err := r.baseRepo.Update("movies", "token", movieToUpdate.Token, updates)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) Inactivate(movieToken string) error {
	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}

	err := r.baseRepo.Update("movies", "token", movieToken, updates)
	if err != nil {
		return err
	}
	return nil
}
