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

func (d *Driver) Create(driver sqlc.InsertDriverParams, createdById int64) (int64, error) {
	driver.FkCreatedByUserID = createdById
	driver.CreatedDate = time.Now()

	return d.db.InsertDriver(context.Background(), driver)
}

func (d *Driver) Update(id int, driver sqlc.UpdateDriverParams, updatedById int64) error {
	driver.ID = int64(id)
	driver.FkUpdatedByUserID = updatedById
	driver.UpdatedDate = time.Now()

	return d.db.UpdateDriver(context.Background(), driver)
}

func (d *Driver) Delete(id int) error {
	return d.db.DeleteDriver(context.Background(), int64(id))
}

func (d *Driver) GetList(search string, offset, limit int) (drivers []sqlc.SelectListDriversRow, total int64, err error) {
	drivers, err = d.db.SelectListDrivers(context.Background(), sqlc.SelectListDriversParams{
		Search: search,
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		return
	}

	total, err = d.db.SelectCountListDrivers(context.Background(), search)

	return
}

func (d *Driver) GetByID(id int) (sqlc.GetDriverRow, error) {
	return d.db.GetDriver(context.Background(), int64(id))
}
