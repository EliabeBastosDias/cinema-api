package threaterrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type ThreaterRepository struct {
	*baserepo.BaseRepository[domain.Threater]
}

var (
	instance *ThreaterRepository
	once     sync.Once
)

func New(db *dbr.Session) *ThreaterRepository {
	once.Do(func() {
		instance = &ThreaterRepository{
			BaseRepository: baserepo.New[domain.Threater](db),
		}
	})
	return instance
}

func (r *ThreaterRepository) Create(threater *domain.Threater) (*domain.Threater, error) {
	threater.Token = uuid.New()
	threater.CreatedAt = time.Now()
	threater.UpdatedAt = time.Now()

	err := r.Insert("threaters", threater)
	if err != nil {
		return nil, err
	}

	return threater, nil
}

func (r *ThreaterRepository) List() ([]domain.Threater, error) {
	var threaters []domain.Threater
	_, err := r.DB().
		Select("*").
		From("threaters").
		Load(&threaters)
	if err != nil {
		return nil, err
	}

	return threaters, nil
}

func (r *ThreaterRepository) Get(threaterToken string) (*domain.Threater, error) {
	var threater domain.Threater

	result, err := r.DB().
		Select("*").
		From("threaters").
		Where("token = ?", threaterToken).
		Load(&threater)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("threater not found")
	}

	return &threater, nil
}

func (r *ThreaterRepository) Update(threaterToUpdate *domain.Threater) error {
	updates := map[string]interface{}{
		"number":      threaterToUpdate.Number,
		"description": threaterToUpdate.Description,
		"updated_at":  time.Now(),
	}

	err := r.BaseRepository.Update("threaters", "token", threaterToUpdate.Token, updates)
	if err != nil {
		return err
	}

	return nil
}

func (r *ThreaterRepository) Inactivate(threaterToken string) error {
	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}

	err := r.BaseRepository.Update("threaters", "token", threaterToken, updates)
	if err != nil {
		return err
	}

	return nil
}
