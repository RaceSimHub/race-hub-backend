package notification_test

import (
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"testing"
	"time"

	mockDb "github.com/RaceSimHub/race-hub-backend/internal/database/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NotificationSuite struct {
	suite.Suite
	notificationService *notification.NotificationImpl
	mockDB              *mockDb.QuerierNotification
}

func (suite *NotificationSuite) SetupTest() {
	suite.mockDB = new(mockDb.QuerierNotification)
	suite.notificationService = notification.NewNotification(suite.mockDB)
}

func (suite *NotificationSuite) TestCreateNotification() {
	suite.mockDB.On("InsertNotification", mock.Anything, mock.AnythingOfType("sqlc.InsertNotificationParams")).Return(int64(1), nil)

	id, err := suite.notificationService.Create("Test message", "Driver1", "Driver2", "Driver3", 10)
	suite.NoError(err)
	suite.Equal(int64(1), id)

	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *NotificationSuite) TestUpdateNotification() {
	suite.mockDB.On("UpdateNotification", mock.Anything, mock.AnythingOfType("sqlc.UpdateNotificationParams")).Return(nil)

	err := suite.notificationService.Update(1, "Test message", "Driver1", "Driver2", "Driver3", 10)
	suite.NoError(err)

	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *NotificationSuite) TestDeleteNotification() {
	suite.mockDB.On("DeleteNotification", mock.Anything, int64(1)).Return(nil)

	err := suite.notificationService.Delete(1)
	suite.NoError(err)

	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *NotificationSuite) TestGetLastMessage() {
	suite.mockDB.On("GetLastNotificationMessage", mock.Anything).Return("Test message", nil)

	message, err := suite.notificationService.GetLastMessage()
	suite.NoError(err)
	suite.Equal("Test message", message)

	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *NotificationSuite) TestGetList() {
	suite.mockDB.On("SelectListNotifications", mock.Anything, mock.AnythingOfType("sqlc.SelectListNotificationsParams")).Return([]sqlc.SelectListNotificationsRow{
		{
			ID:            1,
			Message:       "Test message",
			FirstDriver:   "Driver1",
			SecondDriver:  "Driver2",
			ThirdDriver:   "Driver3",
			LicensePoints: 10,
			CreatedDate:   time.Now(),
		},
	}, nil)

	list, err := suite.notificationService.GetList(0, 10)
	suite.NoError(err)
	suite.Len(list, 1)

	suite.mockDB.AssertExpectations(suite.T())
}

func TestNotificationSuite(t *testing.T) {
	suite.Run(t, new(NotificationSuite))
}
