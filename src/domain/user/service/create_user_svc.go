package service

import (
	"context"
	"database/sql"
	"log"

	"task-management-system/src/domain/user/payload"
	"task-management-system/src/model"
	"task-management-system/src/util"
)

func (s *Service) CreateUserService(
	ctx context.Context,
	request payload.CreateUserRequest,
	userGUID string,
) (err error) {
	password, err := util.GenerateHashPassword(request.Password)
	if err != nil {
		log.Printf("error generate hash password : %s, error : %v", request.Password, err)
		return
	}

	user := model.User{
		GUID:      util.GenerateUUID(),
		Username:  request.Username,
		Email:     request.Email,
		Password:  password,
		RoleGUID:  request.RoleGUID,
		CreatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if err = s.db.Omit("updated_at").Create(&user).Error; err != nil {
		log.Printf("error create user : %v, error : %v", user, err)
		return
	}

	return
}

func (s *Service) CheckEmailIsExistsService(
	ctx context.Context,
	email string,
) (exists bool, err error) {
	var count int64

	if err = s.db.Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		log.Printf("error counting user by email: %s, error: %v", email, err)
		return
	}

	exists = count > 0
	if exists {
		log.Printf("email already exists: %s", email)
	}

	return
}
