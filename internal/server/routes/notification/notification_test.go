package notification_test

import (
	"bytes"
	"encoding/json"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"time"

	mockDb "github.com/RaceSimHub/race-hub-backend/internal/database/mock"
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/notification"
	notificationService "github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NotificationSuite struct {
	suite.Suite
	router            *gin.Engine
	mockService       *notificationService.Notification
	notificationRoute *notification.Notification
	mockDB            *mockDb.QuerierNotification
}

func (suite *NotificationSuite) SetupSuite() {
	suite.mockDB = new(mockDb.QuerierNotification)
	suite.mockService = notificationService.NewNotification(suite.mockDB)
	suite.notificationRoute = notification.NewNotification(*suite.mockService)

	suite.router = gin.Default()
	notificationGroup := suite.router.Group("/notifications")
	{
		notificationGroup.POST("/", suite.notificationRoute.Post)
		notificationGroup.PUT("/:id", suite.notificationRoute.Put)
		notificationGroup.DELETE("/:id", suite.notificationRoute.Delete)
		notificationGroup.GET("/last-message", suite.notificationRoute.GetLastMessage)
		notificationGroup.GET("/list", suite.notificationRoute.GetList)
	}
}

func (suite *NotificationSuite) TestPostNotification() {
	suite.mockDB.On("InsertNotification", mock.Anything, mock.AnythingOfType("sqlc.InsertNotificationParams")).Return(int64(1), nil)

	requestBody := request.PostNotification{
		Message:       "Test message",
		FirstDriver:   "Driver1",
		SecondDriver:  "Driver2",
		ThirdDriver:   "Driver3",
		LicensePoints: 10,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/notifications/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
}

func (suite *NotificationSuite) TestPutNotification() {
	suite.mockDB.On("UpdateNotification", mock.Anything, mock.AnythingOfType("sqlc.UpdateNotificationParams")).Return(nil)

	requestBody := request.PutNotification{
		Message:       "Updated message",
		FirstDriver:   "Driver1",
		SecondDriver:  "Driver2",
		ThirdDriver:   "Driver3",
		LicensePoints: 5,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/notifications/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
}

func (suite *NotificationSuite) TestDeleteNotification() {
	suite.mockDB.On("DeleteNotification", mock.Anything, int64(1)).Return(nil)

	req, _ := http.NewRequest("DELETE", "/notifications/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
}

func (suite *NotificationSuite) TestGetLastMessage() {
	suite.mockDB.On("GetLastNotificationMessage", mock.Anything).Return("Test message", nil)

	req, _ := http.NewRequest("GET", "/notifications/last-message", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
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

	req, _ := http.NewRequest("GET", "/notifications/list?offset=0&limit=10", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func TestNotificationSuite(t *testing.T) {
	suite.Run(t, new(NotificationSuite))
}
