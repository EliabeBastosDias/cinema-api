package mocks

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockThreaterRepository struct {
	mock.Mock
}

func (m *MockThreaterRepository) Get(id string) (*domain.Threater, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Threater), args.Error(1)
}

func (m *MockThreaterRepository) List() ([]domain.Threater, error) {
	args := m.Called()
	return args.Get(0).([]domain.Threater), args.Error(1)
}

func (m *MockThreaterRepository) Create(Threater *domain.Threater) (*domain.Threater, error) {
	args := m.Called(Threater)
	return args.Get(0).(*domain.Threater), args.Error(1)
}

func (m *MockThreaterRepository) Update(Threater *domain.Threater) error {
	args := m.Called(Threater)
	return args.Error(1)
}

func (m *MockThreaterRepository) Inactivate(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
