package user

import (
	"context"
	"errors"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db sqlc.Querier
}

func NewUser(db sqlc.Querier) *User {
	return &User{db: db}
}

func (u *User) Create(email, name, password string) (id int64, err error) {
	password, err = u.hashPassword(password)
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

	isValid := u.checkPasswordHash(password, user.Password)
	if !isValid {
		return "", errors.New("error.invalid.credentials")
	}

	return u.generateToken(int(user.ID))
}

func (u *User) generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(config.JwtSecret)

	return token.SignedString(secret)
}

func (u *User) checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func (u *User) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
