package handlers

import (
	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	patientsComp "github.com/andersonribeir0/ssr-htmx/components/patients"
	"github.com/andersonribeir0/ssr-htmx/domain"
	echo "github.com/labstack/echo/v4"
)

func (a *API) RootHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		patients, err := domain.FindAllPatients(a.db)
		if err != nil {
			c.Error(err)
		}

		patientListComp := patientsComp.PatientList(patients)
		t := templ.Handler(cmpt.Root(patientListComp))

		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
}
