package movietests

import (
	"errors"
	"testing"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMovieCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockMovieRepository)
	mockLogger := new(mocks.MockLogger)
	command := movieservice.NewGetMovieCommand(mockRepo, mockLogger)

	params := movieservice.GetMovieParams{
		MovieToken: "12345",
	}

	expectedMovie := &domain.Movie{
		Token:    uuid.MustParse("12345"),
		Name:     "Inception",
		Director: "Christopher Nolan",
		Duration: 148,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Get", params.MovieToken).Return(expectedMovie, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		result, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedMovie, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("Get", params.MovieToken).Return(nil, errors.New("failed to retrieve movie"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error creating movie: failed to retrieve movie")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
