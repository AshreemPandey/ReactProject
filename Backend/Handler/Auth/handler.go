package Auth

import (
	"Backend/Handler"
	authService "Backend/Service/Auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService *authService.AuthService
}

func (handler *AuthHandler) Login(e echo.Context) error {
	var loginInput Handler.LoginInput
	e.Bind(&loginInput)
	token, err := handler.AuthService.Login(e.Request().Context(), loginInput.Username, loginInput.Password)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return e.JSON(http.StatusOK, map[string]string{"token": token})
}

func InitAuthHandler(authService *authService.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}
