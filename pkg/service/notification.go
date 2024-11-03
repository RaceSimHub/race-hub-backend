package service

import (
	"context"
	"time"

	"github.com/RaceSimHub/race-hub-backend/pkg/database"
	"github.com/RaceSimHub/race-hub-backend/pkg/database/sqlc"
)

type Notification struct{}

func (Notification) Create(message, firstDriver, secondDriver, thirdDriver string, licensePoints int) (int64, error) {
	return database.DbQuerier.InsertNotification(context.Background(), sqlc.InsertNotificationParams{
		Message:       message,
		FirstDriver:   firstDriver,
		SecondDriver:  secondDriver,
		ThirdDriver:   thirdDriver,
		LicensePoints: int32(licensePoints),
		CreatedDate:   time.Now(),
	})
}

func (Notification) Update(id int, message, firstDriver, secondDriver, thirdDriver string, licensePoints int) error {
	return database.DbQuerier.UpdateNotification(context.Background(), sqlc.UpdateNotificationParams{
		ID:            int64(id),
		Message:       message,
		FirstDriver:   firstDriver,
		SecondDriver:  secondDriver,
		ThirdDriver:   thirdDriver,
		LicensePoints: int32(licensePoints),
	})
}

func (Notification) Delete(id int) error {
	return database.DbQuerier.DeleteNotification(context.Background(), int64(id))
}

func (Notification) GetLastMessage() (string, error) {
	return database.DbQuerier.GetLastNotificationMessage(context.Background())
}

func (Notification) GetList(offset, limit int) ([]sqlc.SelectListNotificationsRow, error) {
	return database.DbQuerier.SelectListNotifications(context.Background(), sqlc.SelectListNotificationsParams{
		Column1: int32(offset),
		Column2: int32(limit),
	})
}
