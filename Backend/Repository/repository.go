package repository

import (
	models "Backend/Models"
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Career(ctx context.Context) (*models.Career, error) {
	return &models.Career{
		Title:      "Backend Engineer",
		Position:   "Senior",
		BaseSalary: "300000",
	}, nil
}

func (repo *Repository) FetchUserDetails(ctx context.Context, email string) (*models.LoginCredentials, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)
	return &models.LoginCredentials{
		UserId:   "1",
		UserName: "Ashreem",
		Password: string(password),
	}, nil
}

func (repo *Repository) ShowDashboard(ctx context.Context, username string) (*models.Dashboard, error) {
	return &models.Dashboard{
		UserName: "ashreem",
		FilePath: "./Dashboards/dashboard.pdf",
	}, nil
}

func InitRepository() *Repository {
	return &Repository{
		DB: nil,
	}
}
