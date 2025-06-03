package service

import (
	"context"
	"log"

	"task-management-system/src/handler/jwt"
	"task-management-system/src/model"
)

func (s *Service) LogoutService(
	ctx context.Context,
	claims *jwt.AccessTokenPayload,
) (err error) {
	if err = s.db.Delete(&model.Session{GUID: claims.GUID}).Error; err != nil {
		log.Printf("error delete session by GUID : %s, error: %v", claims.GUID, err.Error())
		return
	}

	return
}
