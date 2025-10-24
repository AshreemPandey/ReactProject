package service

import (
	repository "Backend/Repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	Repository *repository.Repository
}

type Career struct {
	Title    string
	Position string
}

var secretKey = []byte("secret-key")

func (service *Service) Careers(ctx context.Context) (*Career, error) {
	val, err := service.Repository.Career(ctx)
	if err != nil {
		return nil, err
	}
	return &Career{
		Title:    val.Title,
		Position: val.Position,
	}, nil
}

func (service *Service) Login(ctx context.Context, username string, password string) (string, error) {
	credentials, err := service.Repository.Login(ctx, username, password)
	if err != nil {
		return "", err
	}
	if credentials.Username != username || credentials.Password != password {
		return "", errors.New("Invalid login credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return tokenString, nil
}

func (service *Service) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Invalid token")
	}
	return nil
}
