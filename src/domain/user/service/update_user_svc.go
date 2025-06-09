package service

import (
	"context"
	"database/sql"
	"log"

	"task-management-system/src/domain/user/payload"
	"task-management-system/src/model"
	"task-management-system/src/util"
)

func (s *Service) UpdateUserService(
	ctx context.Context,
	request payload.UpdateUserRequest,
	userGUID string,
) (err error) {
	existingUser := model.User{}
	if err = s.db.First(&existingUser, "guid = ?", request.GUID).Error; err != nil {
		log.Printf("error finding user with GUID %s: %v", request.GUID, err)
		return
	}

	if request.Username != "" {
		existingUser.Username = request.Username
	}
	if request.Email != "" {
		existingUser.Email = request.Email
	}
	if request.Password != "" {
		password, errPass := util.GenerateHashPassword(request.Password)
		if errPass != nil {
			log.Printf("error generate hash password : %s, error : %v", request.Password, errPass)
			return errPass
		}
		existingUser.Password = password
	}
	if request.RoleGUID != "" {
		existingUser.RoleGUID = request.RoleGUID
	}

	existingUser.UpdatedBy = sql.NullString{String: userGUID, Valid: true}

	if err = s.db.Save(&existingUser).Error; err != nil {
		log.Printf("error saving updated user: %v", err)
		return
	}

	return
}

func (s *Service) UpdateCheckEmailIsExistsService(
	ctx context.Context,
	guid, email string,
) (exists bool, err error) {
	var count int64

	if err = s.db.Model(&model.User{}).
		Where("email = ? AND guid != ?", email, guid).
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
