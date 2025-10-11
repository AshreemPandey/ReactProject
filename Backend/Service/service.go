package service

import (
	repository "Backend/Repository"
	"context"
)

type Service struct {
	Repository *repository.Repository
}

type Career struct {
	Title    string
	Position string
}

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
