//go:generate mockgen -destination=./mock/driver_mock.go -package=mock . Contract
package driver

import "github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"

type Contract interface {
	Create(name, raceName, email, phone string) (int64, error)
	Update(id int, name, raceName, email, phone string) error
	Delete(id int) error
	GetList(offset, limit int) ([]sqlc.SelectListDriversRow, error)
	GetByID(id int) (sqlc.GetDriverRow, error)
}
