package erp

import (
	repository "Backend/Repository"
	"fmt"

	"github.com/labstack/echo/v4"
)

type ERPService interface {
	ShowDashboard(e echo.Context) error
}

type erpService struct {
	repository *repository.Repository
}

func (service *erpService) ShowDashboard(e echo.Context) error {
	fmt.Println("Welcome Welcome!!!")
	return nil
}

func InitERPService() ERPService {
	return &erpService{
		repository: repository.InitRepository(),
	}
}
