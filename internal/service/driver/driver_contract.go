//go:generate mockgen -destination=./mock/driver_mock.go -package=mock . Contract
package driver

import "github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"

type Contract interface {
	Create(driver sqlc.InsertDriverParams, createdById int64) (int64, error)
	Update(id int, driver sqlc.UpdateDriverParams, updatedById int64) error
	Delete(id int) error
	GetList(offset, limit int) ([]sqlc.SelectListDriversRow, error)
	GetByID(id int) (sqlc.GetDriverRow, error)
}
