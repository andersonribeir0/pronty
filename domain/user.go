package domain

import (
	"strconv"

	"github.com/andersonribeir0/ssr-htmx/components/user"
	"github.com/andersonribeir0/ssr-htmx/db"
)

func FindAllUsers(ds *db.DB) (userList []user.User, err error) {
	users, err := ds.FindAll()
	if err != nil {
		return nil, err
	}

	for _, v := range users {
		userList = append(userList, user.User{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return userList, nil
}

func GetUserDetails(ds *db.DB, id string) string {
	user, err := ds.FindUserByID(id)
	if err != nil {
		return "n/a"
	}

	return strconv.FormatInt(user.DocumentNumber, 10)
}
