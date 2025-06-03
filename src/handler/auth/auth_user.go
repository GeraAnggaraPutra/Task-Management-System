package auth

import (
	"log"

	"task-management-system/src/model"
)

func (a *Auth) User() (data model.User, err error) {
	data = model.User{GUID: a.claims.UserGUID}

	if err = a.db.First(&data).Error; err != nil {
		log.Printf("error find user by GUID : %s, error: %v", a.claims.UserGUID, err)
		return
	}

	return
}
