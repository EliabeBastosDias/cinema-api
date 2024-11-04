package ports

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
)

type SessionRepository interface {
	Create(session *domain.Session) (*domain.Session, error)
	Get(sessionToken string) (*domain.Session, error)
	List(page int) ([]domain.Session, error)
	Update(sessionToUpdate *domain.Session) error
	Inactivate(movieToken string) error
}

type SessionService interface{}
