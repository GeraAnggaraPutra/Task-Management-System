package service

import (
	"context"
	"log"

	"task-management-system/src/domain/user/payload"
	"task-management-system/src/model"
)

func (s *Service) DeleteUserService(
	ctx context.Context,
	request payload.DeleteUserRequest,
) (err error) {
	if err = s.db.Delete(&model.User{GUID: request.GUID}).Error; err != nil {
		log.Printf("error delete user by GUID : %s, error : %v", request.GUID, err)
		return
	}

	return
}
