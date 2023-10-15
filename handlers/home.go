package handlers

import (
	userComp "github.com/andersonribeir0/ssr-htmx/components/user"
	"github.com/andersonribeir0/ssr-htmx/domain"
	"github.com/labstack/echo/v4"
)

func HandleHome(a *API, c echo.Context) error {
	users, err := domain.FindAllUsers(a.db)
	if err != nil {
		c.Error(err)
	}

	userListComp := userComp.UserList(users)

	return PartialHandler(userListComp)(a, c)
}
