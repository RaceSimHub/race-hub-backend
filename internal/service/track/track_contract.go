//go:generate mockgen -destination=./mock/track_mock.go -package=mock . TrackContract
package track

import "github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"

type TrackContract interface {
	Create(name, country string) (int64, error)
	Update(id int, name, country string) error
	Delete(id int) error
	GetList(offset, limit int) ([]sqlc.SelectListTracksRow, error)
	GetByID(id int) (sqlc.SelectTrackByIdRow, error)
}
