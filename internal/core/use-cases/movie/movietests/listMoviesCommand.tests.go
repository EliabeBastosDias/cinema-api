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

func TestListMovieCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockMovieRepository)
	mockLogger := new(mocks.MockLogger)
	command := movieservice.NewListMoviesCommand(mockRepo, mockLogger)

	expectedMovies := []domain.Movie{
		{Name: "Inception", Director: "Christopher Nolan", Duration: 148},
		{Name: "Interstellar", Director: "Christopher Nolan", Duration: 169},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("List").Return(expectedMovies, nil)
		mockLogger.On("Info", mock.Anything).Twice()

		result, err := command.Execute()

		assert.NoError(t, err)
		assert.Equal(t, expectedMovies, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("List").Return(nil, errors.New("database error"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error creating movie: database error")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
