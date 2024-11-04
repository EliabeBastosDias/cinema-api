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

func TestInactivateMovieCommand_Execute(t *testing.T) {
	mockMovieRepo := new(mocks.MockMovieRepository)
	mockSessionRepo := new(mocks.MockSessionRepository)
	mockLogger := new(mocks.MockLogger)
	command := movieservice.NewInactivateMovieCommand(mockMovieRepo, mockSessionRepo, mockLogger)

	params := movieservice.InactivateMovieParams{
		MovieToken: "12345",
	}

	t.Run("Success", func(t *testing.T) {
		mockMovieRepo.On("Get", params.MovieToken).Return(&domain.Movie{}, nil)
		mockMovieRepo.On("Inactivate", params.MovieToken).Return(nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		mockMovieRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Get", func(t *testing.T) {
		mockMovieRepo.On("Get", params.MovieToken).Return(nil, errors.New("movie not found"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockMovieRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Inactivate", func(t *testing.T) {
		mockMovieRepo.On("Get", params.MovieToken).Return(&domain.Movie{}, nil)
		mockMovieRepo.On("Inactivate", params.MovieToken).Return(errors.New("failed to inactivate"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockMovieRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
