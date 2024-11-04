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

func TestUpdateThreaterCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockThreaterRepository)
	mockLogger := new(mocks.MockLogger)
	command := threaterservice.NewUpdateThreaterCommand(mockRepo, mockLogger)

	params := threaterservice.UpdateThreaterParams{
		ThreaterToken: "12345",
		Number:        10,
		Description:   "Updated Description",
	}

	t.Run("Success", func(t *testing.T) {
		existingThreater := &domain.Threater{Token: uuid.MustParse(params.ThreaterToken), Number: 5, Description: "Old Description"}
		mockRepo.On("Get", params.ThreaterToken).Return(existingThreater, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Threater")).Return(nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, params.Number, existingThreater.Number)
		assert.Equal(t, params.Description, existingThreater.Description)
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

	t.Run("Failure on Update", func(t *testing.T) {
		existingThreater := &domain.Threater{Token: uuid.MustParse(params.ThreaterToken)}
		mockRepo.On("Get", params.ThreaterToken).Return(existingThreater, nil)
		mockRepo.On("Update", existingThreater).Return(errors.New("failed to update"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		err := command.Execute(params)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
