package notification

import (
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/stretchr/testify/mock"
)

type MockNotification struct {
	mock.Mock
}

func (m *MockNotification) Create(message, firstDriver, secondDriver, thirdDriver string, licensePoints int) (int64, error) {
	args := m.Called(message, firstDriver, secondDriver, thirdDriver, licensePoints)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockNotification) Update(id int, message, firstDriver, secondDriver, thirdDriver string, licensePoints int) error {
	args := m.Called(id, message, firstDriver, secondDriver, thirdDriver, licensePoints)
	return args.Error(0)
}

func (m *MockNotification) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockNotification) GetLastMessage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockNotification) GetList(offset, limit int) ([]sqlc.SelectListNotificationsRow, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]sqlc.SelectListNotificationsRow), args.Error(1)
}
