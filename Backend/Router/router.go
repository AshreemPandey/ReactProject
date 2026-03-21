package router

import (
	authhandler "Backend/Handler/Auth"
	"Backend/Handler/erp"
	repository "Backend/Repository"
	authService "Backend/Service/Auth"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	authGroup := e.Group("/auth")
	authRouter(authGroup)

	dashboardGroup := e.Group("/dashboard")
	dashboardGroup.Use(authService.AuthMiddleware)
	dashboardRouter(dashboardGroup)

	return e
}

func authRouter(e *echo.Group) {
	authService := authService.InitAuthService(repository.InitRepository())
	handler := authhandler.InitAuthHandler(authService)
	e.POST("/login", handler.Login)
}

func dashboardRouter(e *echo.Group) {
	handler := erp.InitERPHandler()
	e.GET("/", handler.ShowDashboard)
}
