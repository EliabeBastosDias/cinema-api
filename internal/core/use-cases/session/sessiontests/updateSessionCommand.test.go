package sessiontests

import (
	"errors"
	"testing"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	sessionservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/session"
	"github.com/google/uuid"

	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateSessionCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockSessionRepository)
	mockLogger := new(mocks.MockLogger)
	command := sessionservice.NewUpdateSessionCommand(mockRepo, mockLogger)

	params := sessionservice.UpdateSessionParams{
		SessionToken:    "session123",
		MovieToken:      "movie456",
		ThreaterToken:   "threater789",
		SessionDatetime: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		existingSession := &domain.Session{
			Token:           uuid.MustParse(params.SessionToken),
			MovieToken:      "oldMovieToken",
			ThreaterToken:   "oldThreaterToken",
			SessionDatetime: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		}

		mockRepo.On("Get", params.SessionToken).Return(existingSession, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Session")).Return(nil)
		mockLogger.On("Info", mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, params.MovieToken, existingSession.MovieToken)
		assert.Equal(t, params.ThreaterToken, existingSession.ThreaterToken)
		assert.Equal(t, params.SessionDatetime, existingSession.SessionDatetime)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Get", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(nil, errors.New("session not found"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Update", func(t *testing.T) {
		existingSession := &domain.Session{
			Token: uuid.MustParse(params.SessionToken),
		}

		mockRepo.On("Get", params.SessionToken).Return(existingSession, nil)
		mockRepo.On("Update", existingSession).Return(errors.New("failed to update session"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
