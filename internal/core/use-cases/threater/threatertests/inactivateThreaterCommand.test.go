package threatertests

import (
	"errors"
	"testing"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	threaterservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/threater"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInactivateThreaterCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockThreaterRepository)
	mockLogger := new(mocks.MockLogger)
	command := threaterservice.NewInactivateThreaterCommand(mockRepo, nil, mockLogger) // Assume SessionRepository não é necessário neste teste

	params := threaterservice.InactivateThreaterParams{
		ThreaterToken: "12345",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Get", params.ThreaterToken).Return(&domain.Threater{}, nil)
		mockRepo.On("Inactivate", params.ThreaterToken).Return(nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Get", func(t *testing.T) {
		mockRepo.On("Get", params.ThreaterToken).Return(nil, errors.New("threater not found"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure on Inactivate", func(t *testing.T) {
		mockRepo.On("Get", params.ThreaterToken).Return(&domain.Threater{}, nil)
		mockRepo.On("Inactivate", params.ThreaterToken).Return(errors.New("failed to inactivate"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
