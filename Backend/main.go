package main

import (
	"Backend/Router"
	service "Backend/Service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	service := &service.Service{}
	handler := Router.NewHandler(service)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodDelete,
			http.MethodPost,
			http.MethodPut,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAuthorization,
		},
	}))

	e.GET("/careers", handler.GetCareers)
	e.Logger.Fatal(e.Start(":8095"))
}
