package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"task-management-system/src/constant"
	"task-management-system/src/handler/jwt"
)

type Auth struct {
	db     *gorm.DB
	claims *jwt.AccessTokenPayload
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func GetAuth(c *gin.Context) (*Auth, error) {
	a, exists := c.Get("auth")
	if !exists {
		return nil, constant.ErrTokenUnauthorized
	}

	auth, ok := a.(Auth)
	if !ok {
		return nil, constant.ErrTokenUnauthorized
	}

	return &auth, nil
}

func (a *Auth) GetClaims() *jwt.AccessTokenPayload {
	return a.claims
}

func (a *Auth) SetClaims(claims *jwt.AccessTokenPayload) {
	a.claims = claims
}
