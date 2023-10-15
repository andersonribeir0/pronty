package handlers

import (
	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	footerComp "github.com/andersonribeir0/ssr-htmx/components/footer"
	signInComp "github.com/andersonribeir0/ssr-htmx/components/signin"
	"github.com/labstack/echo/v4"
)

func PartialHandler(content templ.Component) func(a *API, c echo.Context) error {
	return func(a *API, c echo.Context) error {
		signIn := signInComp.SignIn()
		footer := footerComp.Footer()
		t := templ.Handler(cmpt.Root(signIn, content, footer))

		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
}
