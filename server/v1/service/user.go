package service

import (
	"Flash_Sale/models"
	daos "Flash_Sale/server/v1/daos/Db"
)

func FindUserById(id int64) (*models.User, error) {
	return daos.SelectUserById(id)
}

func FindUserByName(n string) (*models.User, error) {
	return daos.SelectUserByName(n)
}

func AddNewUser(n string, a int) bool {
	user, err := daos.SelectUserByName(n)
	if err != nil {
		return false
	}
	if user.Name == n {
		return false
	}

	u := &models.User{
		Name: n,
		Age:  a,
	}

	if !daos.InsertUser(u) {
		return false
	}

	return true
}
