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

	e.Static("/static", "./static")

	e.GET("/", h.CreateHandler(handlers.HandleHome))
	e.GET("/user/add", h.CreateHandler(handlers.HandleAddUserForm))
	e.POST("/api/user/add", h.CreateHandler(handlers.HandleAddUser))
	e.DELETE("/api/user", h.CreateHandler(handlers.HandleDeleteUser))
	e.GET("/user/:id/details", h.CreateHandler(handlers.HandleUserDetails))
	e.GET("/api/user/:id", h.CreateHandler(handlers.HandleGetUserItem))

	e.Logger.Fatal(e.Start(os.Getenv("HTTP_LISTEN_ADDR")))
}
