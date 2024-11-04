package sessionrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type SessionRepository struct {
	*baserepo.BaseRepository[domain.Session]
}

var (
	instance *SessionRepository
	once     sync.Once
)

func New(db *dbr.Session) *SessionRepository {
	once.Do(func() {
		instance = &SessionRepository{
			BaseRepository: baserepo.New[domain.Session](db),
		}
	})
	return instance
}

func (r *SessionRepository) Create(session *domain.Session) (*domain.Session, error) {
	session.Token = uuid.New()
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	err := r.Insert("sessions", session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *SessionRepository) List(page int) ([]domain.Session, error) {
	if page < 1 {
		page = 1
	}
	pageSize := 6
	offset := (page - 1) * pageSize

	var sessions []domain.Session
	_, err := r.DB().Select("*").
		From("sessions").
		OrderBy("session_datetime DESC").
		Limit(uint64(pageSize)).
		Offset(uint64(offset)).
		Load(&sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *SessionRepository) Get(sessionToken string) (*domain.Session, error) {
	var session domain.Session

	result, err := r.DB().
		Select("*").
		From("sessions").
		Where("token = ?", sessionToken).
		Load(&session)
	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("session not found")
	}

	return &session, nil
}

func (r *SessionRepository) Update(sessionToUpdate *domain.Session) error {
	updates := map[string]interface{}{
		"session_datetime": sessionToUpdate.SessionDatetime,
		"updated_at":       time.Now(),
	}

	err := r.BaseRepository.Update("sessions", "token", sessionToUpdate.Token, updates)
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Inactivate(sessionToken string) error {
	updates := map[string]interface{}{
		"active":     false,
		"updated_at": time.Now(),
	}

	err := r.BaseRepository.Update("sessions", "token", sessionToken, updates)
	if err != nil {
		return err
	}

	return nil
}
