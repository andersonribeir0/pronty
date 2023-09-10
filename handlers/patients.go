package handlers

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/a-h/templ"
	patientsComp "github.com/andersonribeir0/ssr-htmx/components/patients"
	"github.com/andersonribeir0/ssr-htmx/db"
	"github.com/andersonribeir0/ssr-htmx/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (a *API) HandlerDeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get("hx-target")
		return a.db.DeleteUserById(id)
	}
}

func (a *API) HandleAddUserForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		form := patientsComp.PatientForm()
		t := templ.Handler(form)

		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
}

func (a *API) HandleAddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		docNumberWithoutTokens := removeNonNumeric(c.FormValue("documentNumber"))
		docNumber, err := strconv.ParseInt(docNumberWithoutTokens, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid document number"})
		}

		patient := &db.User{
			ID:             uuid.NewString(),
			Name:           c.FormValue("name"),
			DocumentNumber: docNumber,
			Email:          c.FormValue("email"),
		}

		err = a.db.InsertUser(patient)
		if err != nil {
			return c.Redirect(http.StatusBadGateway, "/")
		}

		patients, err := domain.FindAllPatients(a.db)
		if err != nil {
			c.Error(err)
		}

		patientList := patientsComp.PatientList(patients)

		t := templ.Handler(patientList)
		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
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
