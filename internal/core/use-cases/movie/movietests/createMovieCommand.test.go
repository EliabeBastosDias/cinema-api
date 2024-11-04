package movietests

import (
	"errors"
	"testing"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovieCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockMovieRepository)
	mockLogger := new(mocks.MockLogger)
	command := movieservice.NewCreateMovieCommand(mockRepo, mockLogger)

	params := movieservice.CreateMovieParams{
		Name:     "Inception",
		Director: "Christopher Nolan",
		Duration: 148,
	}

	expectedMovie := &domain.Movie{
		Name:     params.Name,
		Director: params.Director,
		Duration: params.Duration,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*domain.Movie")).Return(expectedMovie, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		result, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedMovie, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*domain.Movie")).Return(nil, errors.New("failed to create movie"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
