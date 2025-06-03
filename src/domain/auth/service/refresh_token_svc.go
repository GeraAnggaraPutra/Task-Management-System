package service

import (
	"context"
	"log"

	"task-management-system/src/domain/auth/helper"
	"task-management-system/src/domain/auth/payload"
	"task-management-system/src/handler/jwt"
	"task-management-system/src/model"

)

func (s *Service) RefreshTokenService(
	ctx context.Context,
	request payload.RefreshTokenRequest,
) (data model.Session, user model.User, err error) {
	refreshTokenClaims, err := jwt.ClaimsRefreshToken(request.RefreshToken)
	if err != nil {
		log.Printf("error claims refresh token : %s, error: %v", request.RefreshToken, err.Error())
		return
	}

	session := model.Session{GUID: refreshTokenClaims.GUID}

	if err = s.db.First(&session).Error; err != nil {
		log.Printf("error find session by GUID : %s, error: %v", refreshTokenClaims.GUID, err.Error())
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(session))
	if err != nil {
		log.Printf("error generate session model : %s, error: %v", request.ToSessionPayload(session), err.Error())
		return
	}

	if err = s.db.Where("guid = ?", session.UserGUID).First(&user).Error; err != nil {
		log.Printf("error find user by GUID : %s, error: %v", session.UserGUID, err.Error())
		return
	}

	if err = s.db.Updates(&data).Error; err != nil {
		log.Printf("error update session : %v, error: %v", data, err.Error())
		return
	}

	return
}
