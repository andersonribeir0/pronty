package handlers

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	usersComp "github.com/andersonribeir0/ssr-htmx/components/user"
	"github.com/andersonribeir0/ssr-htmx/db"
	"github.com/andersonribeir0/ssr-htmx/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func HandleUserDetails(a *API, c echo.Context) error {
	id := c.Param("id")
	details := domain.GetUserDetails(a.db, id)
	t := templ.Handler(usersComp.UserDetails(details))

	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}

func HandleUserList(a *API, c echo.Context) error {
	users, err := domain.FindAllUsers(a.db)
	if err != nil {
		c.Error(err)
	}

	userListComp := usersComp.UserList(users)
	t := templ.Handler(cmpt.Root(userListComp))

	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}

func HandleDeleteUser(a *API, c echo.Context) error {
	id := c.Request().Header.Get("hx-target")
	return a.db.DeleteUserById(id)
}

func HandleAddUserForm(a *API, c echo.Context) error {
	form := usersComp.UserForm("#userRows")
	t := templ.Handler(form)
	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}

func HandleAddUser(a *API, c echo.Context) error {
	docNumberWithoutTokens := removeNonNumeric(c.FormValue("documentNumber"))
	docNumber, err := strconv.ParseInt(docNumberWithoutTokens, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid document number"})
	}

	user := &db.User{
		ID:             uuid.NewString(),
		Name:           c.FormValue("name"),
		DocumentNumber: docNumber,
		Email:          c.FormValue("email"),
	}

	err = a.db.InsertUser(user)
	if err != nil {
		return c.Redirect(http.StatusBadGateway, "/")
	}

	t := templ.Handler(usersComp.UserItem(usersComp.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}))
	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}

func removeNonNumeric(input string) string {
	var result string

	for _, char := range input {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}

	return result
}
