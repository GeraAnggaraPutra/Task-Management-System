package service

import (
	"context"
	"log"

	"task-management-system/src/domain/user/payload"
	"task-management-system/src/model"
)

func (s *Service) ReadUserDetailService(
	ctx context.Context,
	request payload.ReadUserDetailRequest,
) (data model.UserDetail, err error) {
	data = model.UserDetail{GUID: request.GUID}

	if err = s.db.Preload("Role").First(&data).Error; err != nil {
		log.Printf("error find user by GUID : %s, error : %v", request.GUID, err)
		return
	}

	return
}
