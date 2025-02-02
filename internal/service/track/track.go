package track

import (
	"context"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
)

type Track struct {
	db sqlc.Querier
}

func NewTrack(db sqlc.Querier) *Track {
	return &Track{db: db}
}

func (n *Track) Create(name, country string) (int64, error) {
	return n.db.InsertTrack(context.Background(), sqlc.InsertTrackParams{
		Name:    name,
		Country: country,
	})
}

func (n *Track) Update(id int, name, country string) error {
	return n.db.UpdateTrack(context.Background(), sqlc.UpdateTrackParams{
		ID:      int64(id),
		Name:    name,
		Country: country,
	})
}

func (n *Track) Delete(id int) error {
	return n.db.DeleteTrack(context.Background(), int64(id))
}

func (n *Track) GetList(offset, limit int) ([]sqlc.SelectListTracksRow, error) {
	return n.db.SelectListTracks(context.Background(), sqlc.SelectListTracksParams{
		Column1: int32(offset),
		Column2: int32(limit),
	})
}

func (n *Track) GetByID(id int) (sqlc.SelectTrackByIdRow, error) {
	return n.db.SelectTrackById(context.Background(), int64(id))
}
