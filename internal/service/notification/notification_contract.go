package notification

import "github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"

type Notification interface {
	Create(message, firstDriver, secondDriver, thirdDriver string, licensePoints int) (int64, error)
	Update(id int, message, firstDriver, secondDriver, thirdDriver string, licensePoints int) error
	Delete(id int) error
	GetLastMessage() (string, error)
	GetList(offset, limit int) ([]sqlc.SelectListNotificationsRow, error)
}
