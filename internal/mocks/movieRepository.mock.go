package mocks

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) Get(id string) (*domain.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (m *MockMovieRepository) List() ([]domain.Movie, error) {
	args := m.Called()
	return args.Get(0).([]domain.Movie), args.Error(1)
}

func (m *MockMovieRepository) Create(movie *domain.Movie) (*domain.Movie, error) {
	args := m.Called(movie)
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (m *MockMovieRepository) Update(movie *domain.Movie) error {
	args := m.Called(movie)
	return args.Error(1)
}

func (m *MockMovieRepository) Inactivate(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
