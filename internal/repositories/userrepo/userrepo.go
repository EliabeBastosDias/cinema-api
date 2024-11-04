package userrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type UserRepository struct {
	*baserepo.BaseRepository[domain.User]
}

var (
	instance *UserRepository
	once     sync.Once
)

func New(db *dbr.Session) *UserRepository {
	once.Do(func() {
		instance = &UserRepository{
			BaseRepository: baserepo.New[domain.User](db),
		}
	})
	return instance
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	user.Token = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.LastLogin = time.Now()
	user.Active = true

	err := r.Insert("users", user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Get(userToken string) (*domain.User, error) {
	var user domain.User

	result, err := r.DB().
		Select("*").
		From("users").
		Where("token = ?", userToken).
		Load(&user)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (r *UserRepository) List() ([]domain.User, error) {
	var users []domain.User

	_, err := r.DB().
		Select("*").
		From("users").
		Load(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(userToUpdate *domain.User) error {
	updates := map[string]interface{}{
		"username":   userToUpdate.Username,
		"email":      userToUpdate.Email,
		"password":   userToUpdate.Password,
		"role":       userToUpdate.Roles,
		"updated_at": time.Now(),
	}

	err := r.BaseRepository.Update("users", "token", userToUpdate.Token, updates)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Inactivate(userToken string) error {
	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}

	err := r.BaseRepository.Update("users", "token", userToken, updates)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	result, err := r.DB().
		Select("*").
		From("users").
		Where("email = ?", email).
		Load(&user)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}
