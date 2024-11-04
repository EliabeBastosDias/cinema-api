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

func TestListSessionsCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockSessionRepository)
	mockLogger := new(mocks.MockLogger)
	command := sessionservice.NewListSessionsCommand(mockRepo, mockLogger)

	params := sessionservice.ListSessionsParams{
		Page: 1,
	}

	expectedSessions := []domain.Session{
		{Token: uuid.MustParse("1"), MovieToken: "movie1", ThreaterToken: "threater1", SessionDatetime: time.Now()},
		{Token: uuid.MustParse("2"), MovieToken: "movie2", ThreaterToken: "threater2", SessionDatetime: time.Now()},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("List", params.Page).Return(expectedSessions, nil)
		mockLogger.On("Info", mock.Anything).Twice()

		sessions, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedSessions, sessions)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("List", params.Page).Return(nil, errors.New("database error"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		sessions, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, sessions)
		assert.EqualError(t, err, "error creating Session: database error")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
