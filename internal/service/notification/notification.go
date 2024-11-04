package notification

import (
	"context"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
)

type NotificationImpl struct {
	db sqlc.Querier
}

func NewNotification(db sqlc.Querier) *NotificationImpl {
	return &NotificationImpl{db: db}
}

func (n *NotificationImpl) Create(message, firstDriver, secondDriver, thirdDriver string, licensePoints int) (int64, error) {
	return n.db.InsertNotification(context.Background(), sqlc.InsertNotificationParams{
		Message:       message,
		FirstDriver:   firstDriver,
		SecondDriver:  secondDriver,
		ThirdDriver:   thirdDriver,
		LicensePoints: int32(licensePoints),
		CreatedDate:   time.Now(),
	})
}

func (n *NotificationImpl) Update(id int, message, firstDriver, secondDriver, thirdDriver string, licensePoints int) error {
	return n.db.UpdateNotification(context.Background(), sqlc.UpdateNotificationParams{
		ID:            int64(id),
		Message:       message,
		FirstDriver:   firstDriver,
		SecondDriver:  secondDriver,
		ThirdDriver:   thirdDriver,
		LicensePoints: int32(licensePoints),
	})
}

func (n *NotificationImpl) Delete(id int) error {
	return n.db.DeleteNotification(context.Background(), int64(id))
}

func (n *NotificationImpl) GetLastMessage() (string, error) {
	return n.db.GetLastNotificationMessage(context.Background())
}

func (n *NotificationImpl) GetList(offset, limit int) ([]sqlc.SelectListNotificationsRow, error) {
	return n.db.SelectListNotifications(context.Background(), sqlc.SelectListNotificationsParams{
		Column1: int32(offset),
		Column2: int32(limit),
	})
}
