package helper

import (
	"context"
	"log"

	"task-management-system/src/domain/auth/payload"
	"task-management-system/src/handler/jwt"
	"task-management-system/src/model"
)

func GenerateSessionModel(
	ctx context.Context,
	request payload.SessionPayload,
) (data model.Session, err error) {
	accessToken, err := jwt.GenerateAccessToken(request.ToAccessTokenRequest())
	if err != nil {
		log.Printf("error generate access token, error: %v", err)
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(request.ToRefreshTokenRequest())
	if err != nil {
		log.Printf("error generate refresh token, error: %v", err)
		return
	}

	data = model.Session{
		GUID:                  request.SessionGUID,
		UserGUID:              request.UserGUID,
		AccessToken:           accessToken.Token,
		AccessTokenExpiresAt:  accessToken.ExpiresAt,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiresAt: refreshToken.ExpiresAt,
	}

	return
}
