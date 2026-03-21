package Handler

import (
	service "Backend/Service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ERPHandler interface {
	ShowDashboard(e echo.Context) error
}

type erpHandler struct {
	ServiceLayer *service.Service
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewHandler(service *service.Service) ERPHandler {
	return &erpHandler{
		ServiceLayer: service,
	}
}

// Hello godoc
// @Tags GetCareers
// @Summary Get Careers
// @Description Returns list of careers
// @Produce  json
// @Success 200 {object} service.Career
// @Router /careers [get]
func (handler *erpHandler) GetCareers(e echo.Context) error {
	//val,err:=service
	fmt.Println("In GetCareers")
	ctx := e.Request().Context()
	careers, err := handler.ServiceLayer.Careers(ctx)
	if err != nil {
		return e.JSON(500, "Internal Server Error")
	}
	return e.JSON(http.StatusOK, careers)
}

// Hello godoc
// @Tags Login
// @Summary User Login
// @Description Allows user to login
// @Accept  json
// @Produce  json
// @Param login body LoginRequest true "Login Credentials"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (handler *erpHandler) Login(e echo.Context) error {
	fmt.Println("In Login Handler")
	var req LoginRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(400, "Bad Request")
	}
	ctx := e.Request().Context()
	username := req.Username
	password := req.Password

	token, err := handler.ServiceLayer.Login(ctx, username, password)
	if err != nil {
		return e.JSON(401, "Unauthorized")
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (handler *erpHandler) ShowDashboard(e echo.Context) error {
	_, err := handler.ServiceLayer.ShowDashboard(e.Request().Context(), e.Request().Context().Value("user-id").(string))
	if err != nil {
		return e.JSON(500, "Internal Server Error")
	}
	return e.JSON(http.StatusOK, "Dashboard")
}
