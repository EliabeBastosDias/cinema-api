package threatertests

import (
	"errors"
	"testing"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	threaterservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/threater"
	"github.com/EliabeBastosDias/cinema-api/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetThreaterCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockThreaterRepository)
	mockLogger := new(mocks.MockLogger)
	command := threaterservice.NewGetThreaterCommand(mockRepo, mockLogger)

	params := threaterservice.GetThreaterParams{
		ThreaterToken: "12345",
	}

	expectedThreater := &domain.Threater{
		Token:       uuid.MustParse(params.ThreaterToken),
		Number:      1,
		Description: "A large threater with many seats",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Get", params.ThreaterToken).Return(expectedThreater, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		result, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedThreater, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("Get", params.ThreaterToken).Return(nil, errors.New("database error"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error creating threater: database error")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
