package repository

import (
	models "Backend/Models"
	"context"
)

type Repository struct {
}

func (repo *Repository) Career(ctx context.Context) (*models.Career, error) {
	return &models.Career{
		Title:      "Backend Engineer",
		Position:   "Senior",
		BaseSalary: "300000",
	}, nil
}

func (repo *Repository) Login(ctx context.Context, username string, password string) (*models.LoginCredentials, error) {
	return &models.LoginCredentials{
		Username: "ashreem",
		Password: "password123",
	}, nil
}
