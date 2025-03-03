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

func (n *Track) GetList(search string, offset, limit int) (tracks []sqlc.SelectListTracksRow, total int64, err error) {
	tracks, err = n.db.SelectListTracks(context.Background(), sqlc.SelectListTracksParams{
		Search:  search,
		Column1: int32(offset),
		Column2: int32(limit),
	})

	if err != nil {
		return
	}

	total, err = n.db.SelectListTracksCount(context.Background(), search)

	return
}

func (n *Track) GetByID(id int) (sqlc.SelectTrackByIdRow, error) {
	return n.db.SelectTrackById(context.Background(), int64(id))
}
