package mocks

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockSessionRepository struct {
	mock.Mock
}

func (m *MockSessionRepository) Get(id string) (*domain.Session, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Session), args.Error(1)
}

func (m *MockSessionRepository) List(page int) ([]domain.Session, error) {
	args := m.Called()
	return args.Get(0).([]domain.Session), args.Error(1)
}

func (m *MockSessionRepository) Create(Session *domain.Session) (*domain.Session, error) {
	args := m.Called(Session)
	return args.Get(0).(*domain.Session), args.Error(1)
}

func (m *MockSessionRepository) Update(Session *domain.Session) error {
	args := m.Called(Session)
	return args.Error(1)
}

func (m *MockSessionRepository) Inactivate(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
