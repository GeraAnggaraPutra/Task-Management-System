package service

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"task-management-system/src/constant"
	"task-management-system/src/domain/auth/helper"
	"task-management-system/src/domain/auth/payload"
	"task-management-system/src/model"
	"task-management-system/src/util"
)

func (s *Service) LoginService(
	ctx context.Context,
	request payload.LoginRequest,
) (data model.Session, user model.User, err error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		log.Printf("error begin transaction, error: %v", tx.Error)
		err = errors.WithStack(constant.ErrUnknownSource)
		return
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback().Error; errRollback != nil {
				log.Printf("error rollback transaction, error: %v", errRollback)
			}
		}
	}()

	if err = s.db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		log.Printf("error find user by email : %s, error: %v", request.Email, err)
		err = constant.ErrUserNotFound
		return
	}

	if err = util.CompareHashPassword(request.Password, user.Password); err != nil {
		log.Printf("password incorrect for user : %s, error: %v", user.Email, err)
		err = constant.ErrPasswordIncorrect
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(user.GUID))
	if err != nil {
		log.Printf("error generate session model : %s, error: %v", request.ToSessionPayload(user.GUID), err)
		return
	}

	if err = tx.Create(&data).Error; err != nil {
		log.Printf("error create session : %v, error: %v", data, err)
		return
	}

	if err = tx.Commit().Error; err != nil {
		log.Printf("error commit transaction, error: %v", err)
		err = errors.WithStack(constant.ErrUnknownSource)
		return
	}

	return
}
