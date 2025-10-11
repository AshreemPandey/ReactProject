package Router

import (
	service "Backend/Service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	ServiceLayer *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		ServiceLayer: service,
	}
}

func (handler *Handler) GetCareers(e echo.Context) error {
	//val,err:=service
	fmt.Println("In GetCareers")
	ctx := e.Request().Context()
	careers, err := handler.ServiceLayer.Careers(ctx)
	if err != nil {
		return e.JSON(500, "Internal Server Error")
	}
	return e.JSON(http.StatusOK, careers)
}
