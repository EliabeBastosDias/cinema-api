package sessiontests

import (
	"errors"
	"testing"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	sessionservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/session"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInactivateSessionCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockSessionRepository)
	mockLogger := new(mocks.MockLogger)
	command := sessionservice.NewInactivateSessionCommand(mockRepo, mockLogger)

	params := sessionservice.InactivateSessionParams{
		SessionToken: "test-session-token",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(&domain.Session{}, nil)
		mockRepo.On("Inactivate", params.SessionToken).Return(nil)
		mockLogger.On("Info", mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Get", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(nil, errors.New("session not found"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		assert.EqualError(t, err, "session not found")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Inactivate", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(&domain.Session{}, nil)
		mockRepo.On("Inactivate", params.SessionToken).Return(errors.New("inactivation failed"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		assert.EqualError(t, err, "inactivation failed")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
