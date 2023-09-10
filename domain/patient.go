package domain

import (
	"github.com/andersonribeir0/ssr-htmx/components/patients"
	"github.com/andersonribeir0/ssr-htmx/db"
)

func FindAllPatients(ds *db.DB) (patientList []patients.Patient, err error) {
	users, err := ds.FindAll()
	if err != nil {
		return nil, err
	}

	for _, v := range users {
		patientList = append(patientList, patients.Patient{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return patientList, nil
}
