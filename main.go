package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/andersonribeir0/ssr-htmx/components/patients"
	patientsComp "github.com/andersonribeir0/ssr-htmx/components/patients"
)

func main() {
	patients := []patients.Patient{
		{
			Name:  "John Jimenez",
			Email: "jj@kkj.com",
		},
		{
			Name:  "Lucy Lil",
			Email: "ll@ll.com",
		},
	}

	patientListComp := patientsComp.PatientList(patients)
	http.Handle("/", templ.Handler(root(patientListComp)))

	http.ListenAndServe(":8080", nil)
}
