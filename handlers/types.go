package handlers

import (
	"os"

	"github.com/andersonribeir0/ssr-htmx/db"
	database "github.com/andersonribeir0/ssr-htmx/db"
	echo "github.com/labstack/echo/v4"
)

type API struct {
	db *db.DB
}

func NewAPI() *API {
	db := database.NewDB(os.Getenv("DATABASE_URL"))
	if err := db.Ping(); err != nil {
		panic("database connection failed: " + err.Error())
	}

	return &API{
		db: db,
	}
}

func (a *API) CreateHandler(handlerFunc func(*API, echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handlerFunc(a, c)
	}
}
