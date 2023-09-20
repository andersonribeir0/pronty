package handlers

import (
	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	userComp "github.com/andersonribeir0/ssr-htmx/components/user"
	"github.com/andersonribeir0/ssr-htmx/domain"
	echo "github.com/labstack/echo/v4"
)

func (a *API) RootHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := domain.FindAllUsers(a.db)
		if err != nil {
			c.Error(err)
		}

		userListComp := userComp.UserList(users)
		t := templ.Handler(cmpt.Root(userListComp))

		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
}
