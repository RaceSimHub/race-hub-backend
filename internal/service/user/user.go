package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/model"
	"github.com/RaceSimHub/race-hub-backend/pkg/email"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db           sqlc.Querier
	serviceEmail *email.Email
}

func NewUser(db sqlc.Querier) *User {
	serviceEmail := email.NewEmail(config.RaceHubHost, config.EmailFrom, config.EmailPassword, config.EmailHost, config.EmailPort)

	return &User{
		db:           db,
		serviceEmail: serviceEmail,
	}
}

func (u *User) Create(email, name, password string) (id int64, err error) {
	password, err = u.hashPassword(password)
	if err != nil {
		return
	}
	emailVerificationToken := uuid.New().String()
	emailVerificationToken = emailVerificationToken[:8]

	id, err = u.db.InsertUser(context.Background(), sqlc.InsertUserParams{
		Email:                      email,
		Name:                       name,
		Password:                   password,
		Status:                     string(model.UserStatusPending),
		EmailVerificationToken:     emailVerificationToken,
		EmailVerificationExpiresAt: time.Now().Add(time.Hour * 24),
		Role:                       string(model.UserRoleDriver),
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, errors.New("email já cadastrado")
		}

		return
	}

	u.serviceEmail.SendUserCreatedEmail(email, name, emailVerificationToken)

	return
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

	if model.UserStatus(user.Status) != model.UserStatusActive {
		return "", errors.New("error.user.not.active")
	}

	return u.generateToken(int(user.ID))
}

func (u *User) VerifyEmail(email, token string) (err error) {
	user, err := u.db.SelectUserByEmailVerificationToken(context.Background(), sqlc.SelectUserByEmailVerificationTokenParams{
		EmailVerificationToken: token,
		Email:                  email,
		Status:                 string(model.UserStatusPending),
	})
	if err != nil {
		return
	}

	if user.EmailVerificationExpiresAt.Before(time.Now()) {
		return errors.New("error.email.verification.expired")
	}

	err = u.db.UpdateUserStatus(context.Background(), sqlc.UpdateUserStatusParams{
		ID:     user.ID,
		Status: string(model.UserStatusActive),
	})

	return
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
