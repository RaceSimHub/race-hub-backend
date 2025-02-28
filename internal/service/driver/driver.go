package driver

import (
	"context"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
)

type Driver struct {
	db sqlc.Querier
}

func NewDriver(db sqlc.Querier) *Driver {
	return &Driver{db: db}
}

func (d *Driver) Create(name, raceName, email, phone string) (int64, error) {
	return d.db.InsertDriver(context.Background(), sqlc.InsertDriverParams{
		Name:              name,
		RaceName:          raceName,
		Email:             email,
		Phone:             phone,
		FkCreatedByUserID: 1,
	})
}

func (d *Driver) Update(id int, name, raceName, email, phone string) error {
	return d.db.UpdateDriver(context.Background(), sqlc.UpdateDriverParams{
		ID:                int64(id),
		Name:              name,
		RaceName:          raceName,
		Email:             email,
		Phone:             phone,
		FkUpdatedByUserID: 1,
		UpdatedDate:       time.Now(),
	})
}

func (d *Driver) Delete(id int) error {
	return d.db.DeleteDriver(context.Background(), int64(id))
}

func (d *Driver) GetList(offset, limit int) ([]sqlc.SelectListDriversRow, error) {
	return d.db.SelectListDrivers(context.Background(), sqlc.SelectListDriversParams{
		Column1: int32(offset),
		Column2: int32(limit),
	})
}

func (d *Driver) GetByID(id int) (sqlc.GetDriverRow, error) {
	return d.db.GetDriver(context.Background(), int64(id))
}
