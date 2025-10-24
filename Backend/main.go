package main

import (
	"Backend/Router"
	service "Backend/Service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "Backend/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title My API
// @version 1.0
// @description This is the API documentation for my backend
// @host localhost:8095
// @BasePath /
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
			echo.HeaderContentType,
		},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/careers", handler.GetCareers)
	e.POST("/login", handler.Login)
	e.Logger.Fatal(e.Start(":8095"))
}
