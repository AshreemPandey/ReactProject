package auth

import (
	repository "Backend/Repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

type AuthService struct {
	Repository *repository.Repository
}

func (service *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	credentials, err := service.Repository.FetchUserDetails(ctx, email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid login credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": credentials.UserId,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return tokenString, nil
}

func (service *AuthService) VerifyToken(tokenString string) error {
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

func InitAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		Repository: repo,
	}
}
