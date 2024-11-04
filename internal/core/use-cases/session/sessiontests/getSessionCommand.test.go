package sessiontests

import (
	"errors"
	"testing"
	"time"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	sessionservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/session"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSessionCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockSessionRepository)
	mockLogger := new(mocks.MockLogger)
	command := sessionservice.NewGetSessionCommand(mockRepo, mockLogger)

	params := sessionservice.GetSessionParams{
		SessionToken: "test-token",
	}

	expectedSession := &domain.Session{
		Token:           uuid.MustParse(params.SessionToken),
		MovieToken:      "movie123",
		ThreaterToken:   "threater123",
		SessionDatetime: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(expectedSession, nil)
		mockLogger.On("Info", mock.Anything).Twice()

		result, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedSession, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("Get", params.SessionToken).Return(nil, errors.New("session not found"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error creating session: session not found")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
