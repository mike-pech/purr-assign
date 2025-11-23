package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mike-pech/purr-assign/cmd/api/v1"
	v1 "github.com/mike-pech/purr-assign/internal/delivery/http/v1"
	"github.com/mike-pech/purr-assign/internal/repository"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

//go:embed migrations/*.sql
var sqlMigrations embed.FS
var (
	validate   *validator.Validate
	migrations = migrate.NewMigrations()
)

type Validator struct {
	v *validator.Validate
}

// Подогнан под структуры данных
func (v Validator) Validate(i any) error {
	return v.v.Struct(i)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	v := Validator{v: validate}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)),
	))

	err = migrations.Discover(sqlMigrations)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewBunRepository(sqldb)

	server := v1.NewServer(repo, v)

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
