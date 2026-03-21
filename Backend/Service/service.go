package service

import (
	models "Backend/Models"
	repository "Backend/Repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	credentials, err := service.Repository.FetchUserDetails(ctx, username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid login credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return tokenString, nil
}

func (service *Service) ShowDashboard(ctx context.Context, username string) (*models.Dashboard, error) {
	return service.Repository.ShowDashboard(ctx, username)
}

func (service *Service) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return errors.New("Failed parsing token")
	}
	if !token.Valid {
		return errors.New("Invalid token")
	}
	return nil
}
