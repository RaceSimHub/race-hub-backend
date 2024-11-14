package mock

import (
	"context"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/stretchr/testify/mock"
)

type QuerierNotification struct {
	mock.Mock
}

func (m *QuerierNotification) InsertUser(ctx context.Context, arg sqlc.InsertUserParams) (int64, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(int64), args.Error(1)
}

func (m *QuerierNotification) SelectUserIDByEmailAndPassword(ctx context.Context, email string) (int64, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(int64), args.Error(1)
}

func (m *QuerierNotification) InsertNotification(ctx context.Context, arg sqlc.InsertNotificationParams) (int64, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(int64), args.Error(1)
}

func (m *QuerierNotification) UpdateNotification(ctx context.Context, arg sqlc.UpdateNotificationParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *QuerierNotification) DeleteNotification(ctx context.Context, arg int64) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *QuerierNotification) GetLastNotificationMessage(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *QuerierNotification) SelectListNotifications(ctx context.Context, arg sqlc.SelectListNotificationsParams) ([]sqlc.SelectListNotificationsRow, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).([]sqlc.SelectListNotificationsRow), args.Error(1)
}
