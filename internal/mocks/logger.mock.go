package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Warn(message string, err error) {
	m.Called(message, err)
}

func (m *MockLogger) DPanic(recoveryData any, msg ...string) {
	m.Called(recoveryData, msg)
}

func (m *MockLogger) Panic(auditUuid string, recoveryData any) {
	m.Called(auditUuid, recoveryData)
}

func (m *MockLogger) Error(message string, err error) {
	m.Called(message, err)
}

func (m *MockLogger) Info(message string, ctx ...interface{}) {
	m.Called(message, ctx)
}
