package erp

import (
	"Backend/Service/erp"

	"github.com/labstack/echo/v4"
)

type ERPHandler interface {
	ShowDashboard(e echo.Context) error
}

type erpHandler struct {
	service erp.ERPService
}

func (handler *erpHandler) ShowDashboard(e echo.Context) error {
	handler.service.ShowDashboard(e)
	return nil
}

func InitERPHandler() ERPHandler {
	return &erpHandler{
		service: erp.InitERPService(),
	}
}
