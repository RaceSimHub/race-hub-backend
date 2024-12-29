package notification_test

import (
	"testing"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"go.uber.org/mock/gomock"

	mockDb "github.com/RaceSimHub/race-hub-backend/internal/database/mock"
	"github.com/stretchr/testify/suite"
)

type NotificationSuite struct {
	suite.Suite
	notificationService *notification.Notification
	mockDB              *mockDb.MockQuerier
}

func (suite *NotificationSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockDB = mockDb.NewMockQuerier(ctrl)
	suite.notificationService = notification.NewNotification(suite.mockDB)
}

func (suite *NotificationSuite) TestCreateNotification() {
	suite.mockDB.EXPECT().InsertNotification(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	id, err := suite.notificationService.Create("Test message", "Driver1", "Driver2", "Driver3", 10)
	suite.NoError(err)
	suite.Equal(int64(1), id)
}

func (suite *NotificationSuite) TestUpdateNotification() {
	suite.mockDB.EXPECT().UpdateNotification(gomock.Any(), gomock.Any()).Return(nil)

	err := suite.notificationService.Update(1, "Test message", "Driver1", "Driver2", "Driver3", 10)
	suite.NoError(err)
}

func (suite *NotificationSuite) TestDeleteNotification() {
	suite.mockDB.EXPECT().DeleteNotification(gomock.Any(), int64(1)).Return(nil)

	err := suite.notificationService.Delete(1)
	suite.NoError(err)
}

func (suite *NotificationSuite) TestGetLastMessage() {
	suite.mockDB.EXPECT().GetLastNotificationMessage(gomock.Any()).Return("Test message", nil)

	message, err := suite.notificationService.GetLastMessage()
	suite.NoError(err)
	suite.Equal("Test message", message)
}

func (suite *NotificationSuite) TestGetList() {
	suite.mockDB.EXPECT().SelectListNotifications(gomock.Any(), gomock.Any()).Return([]sqlc.SelectListNotificationsRow{
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
}

func TestNotificationSuite(t *testing.T) {
	suite.Run(t, new(NotificationSuite))
}
