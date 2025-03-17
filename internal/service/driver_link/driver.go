package driverlink

import (
	"context"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/model"
)

type DriverLink struct {
	db sqlc.Querier
}

func NewDriverLink(db sqlc.Querier) *DriverLink {
	return &DriverLink{db: db}
}

func (d *DriverLink) Create(driverID, userID int64) (int64, error) {
	insertDriverLinkParams := sqlc.InsertDriverLinkParams{
		FkDriverID: driverID,
		FkUserID:   userID,
		Status:     string(model.DriverLinkStatusPending),
		CreatedAt:  time.Now(),
	}

	return d.db.InsertDriverLink(context.Background(), insertDriverLinkParams)
}

func (d *DriverLink) GetList(search string, offset, limit int) (driverLinks []sqlc.SelectDriverLinksRow, total int64, err error) {
	driverLinks, err = d.db.SelectDriverLinks(context.Background(), sqlc.SelectDriverLinksParams{
		Search: search,
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		return
	}

	total, err = d.db.SelectCountDriverLinks(context.Background(), search)

	return
}

func (d *DriverLink) GetStatusByUserID(userID int64) (status string, err error) {
	return d.db.SelectDriverLinkStatusByUserID(context.Background(), userID)
}
