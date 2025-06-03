package auth

import (
	"context"
	"log"
	"time"

	"task-management-system/src/constant"
	"task-management-system/src/model"
)

func (a *Auth) ValidateSession(ctx context.Context) (err error) {
	var session model.Session

	result := a.db.WithContext(ctx).
		Where("guid = ?", a.claims.GUID).
		First(&session)

	if result.Error != nil {
		log.Printf("error read session by GUID : %s, error: %v", a.claims.GUID, result.Error)
		return result.Error
	}

	if time.Now().After(session.AccessTokenExpiresAt) {
		err = constant.ErrTokenExpired
		return
	}

	return
}
