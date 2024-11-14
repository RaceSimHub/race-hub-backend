package user

import (
	"context"
	"errors"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
)

type User struct {
	db sqlc.Querier
}

func NewUser(db sqlc.Querier) *User {
	return &User{db: db}
}

func (u *User) Create(email, name, password string) (id int64, err error) {
	password, err = utils.HashPassword(password)
	if err != nil {
		return
	}

	return u.db.InsertUser(context.Background(), sqlc.InsertUserParams{
		Email:    email,
		Name:     name,
		Password: password,
	})
}

func (u *User) GenerateToken(email, password string) (token string, err error) {
	user, err := u.db.SelectUserIDAndPasswordByEmail(context.Background(), email)
	if err != nil {
		return
	}

	isValid := utils.CheckPasswordHash(password, user.Password)
	if !isValid {
		return "", errors.New("error.invalid.credentials")
	}

	return utils.GenerateToken(int(user.ID))
}
