package handlers

import (
	"github.com/a-h/templ"
	cmpt "github.com/andersonribeir0/ssr-htmx/components"
	patientsComp "github.com/andersonribeir0/ssr-htmx/components/patients"
	echo "github.com/labstack/echo/v4"
)

func (a *API) RootHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var patients []patientsComp.Patient

		users, err := a.db.FindAll()
		if err != nil {
			c.Logger().Error(err)
			return c.NoContent(echo.ErrNotFound.Code)
		}

		for _, v := range users {
			patients = append(patients, patientsComp.Patient{
				Name:  v.Name,
				Email: v.Email,
			})
		}

		patientListComp := patientsComp.PatientList(patients)
		t := templ.Handler(cmpt.Root(patientListComp))

		return t.Component.Render(c.Request().Context(), c.Response().Writer)
	}
}
