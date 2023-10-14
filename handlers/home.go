package handlers

import (
	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	footerComp "github.com/andersonribeir0/ssr-htmx/components/footer"
	signInComp "github.com/andersonribeir0/ssr-htmx/components/signin"
	userComp "github.com/andersonribeir0/ssr-htmx/components/user"
	"github.com/andersonribeir0/ssr-htmx/domain"
	"github.com/labstack/echo/v4"
)

func HandleHome(a *API, c echo.Context) error {
	users, err := domain.FindAllUsers(a.db)
	if err != nil {
		c.Error(err)
	}

	signIn := signInComp.SignIn()
	userListComp := userComp.UserList(users)
	footer := footerComp.Footer()

	t := templ.Handler(cmpt.Root(signIn, userListComp, footer))

	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}
