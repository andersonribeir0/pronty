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
	e.GET("/user/add", h.HandleAddUserForm())
	e.POST("/api/user/add", h.HandleAddUser())
	e.DELETE("/api/user", h.HandlerDeleteUser())
	e.Logger.Fatal(e.Start(os.Getenv("HTTP_LISTEN_ADDR")))
}
