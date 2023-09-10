package main

import (
	"log"
	"os"

	"github.com/andersonribeir0/ssr-htmx/handlers"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	h := handlers.NewAPI()

	e.GET("/", h.RootHandler())
	e.GET("/patient/add", h.HandleAddUserForm())
	e.POST("/patient/add", h.HandleAddUser())
	e.DELETE("/patient", h.HandlerDeleteUser())

	e.Logger.Fatal(e.Start(os.Getenv("HTTP_LISTEN_ADDR")))
}
