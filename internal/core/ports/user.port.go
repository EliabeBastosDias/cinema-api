package ports

import "github.com/EliabeBastosDias/cinema-api/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Get(userToken string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]domain.User, error)
	Update(userToUpdate *domain.User) error
	Inactivate(userToken string) error
}

type UserService interface{}
