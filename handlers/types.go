package handlers

import (
	"os"

	"github.com/andersonribeir0/ssr-htmx/db"
	database "github.com/andersonribeir0/ssr-htmx/db"
)

type API struct {
	db *db.DB
}

func NewAPI() *API {
	return &API{
		db: database.NewDB(os.Getenv("DATABASE_URL")),
	}
}
