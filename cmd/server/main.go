package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mike-pech/purr-assign/cmd/api"
)

var validate *validator.Validate

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func main() {
	server := NewServer()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	api.RegisterHandlers(e, server)

	e.GET("/swagger/*", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, swagger)
	})
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return ctx.Redirect(http.StatusMovedPermanently, "/swagger/doc.json")
		}
	})

	log.Fatal(e.Start("0.0.0.0:8080"))
}
