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

func TestUpdateMovieCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockMovieRepository)
	mockLogger := new(mocks.MockLogger)
	command := movieservice.NewUpdateMovieCommand(mockRepo, mockLogger)

	params := movieservice.UpdateMovieParams{
		MovieToken: "12345",
		Name:       "Updated Name",
		Director:   "Updated Director",
		Duration:   150,
	}

	t.Run("Success", func(t *testing.T) {
		existingMovie := &domain.Movie{Token: uuid.MustParse(params.MovieToken), Name: "Old Name", Director: "Old Director", Duration: 120}
		mockRepo.On("Get", params.MovieToken).Return(existingMovie, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Movie")).Return(nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, params.Name, existingMovie.Name)
		assert.Equal(t, params.Director, existingMovie.Director)
		assert.Equal(t, params.Duration, existingMovie.Duration)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Get", func(t *testing.T) {
		mockRepo.On("Get", params.MovieToken).Return(nil, errors.New("movie not found"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Update", func(t *testing.T) {
		existingMovie := &domain.Movie{Token: uuid.MustParse(params.MovieToken)}
		mockRepo.On("Get", params.MovieToken).Return(existingMovie, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Movie")).Return(errors.New("failed to update"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
