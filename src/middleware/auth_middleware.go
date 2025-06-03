package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"task-management-system/src/constant"
	"task-management-system/src/handler/auth"
	"task-management-system/src/handler/jwt"
)

type EnsureToken struct {
	auth *auth.Auth
}

func NewEnsureToken(db *gorm.DB) *EnsureToken {
	ah := auth.NewAuth(db)

	return &EnsureToken{
		auth: ah,
	}
}

func (et *EnsureToken) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := parseHeaderToken(c.Request.Header)
		if err != nil {
			log.Printf("Error parse header token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessTokenClaims, err := jwt.ClaimsAccessToken(token)
		if err != nil {
			log.Printf("Error claims access token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		et.auth.SetClaims(&accessTokenClaims)

		err = et.auth.ValidateSession(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("auth", *et.auth)
		c.Next()
	}
}

func parseHeaderToken(h http.Header) (token string, err error) {
	headerDataToken := h.Get(constant.DefaultMdwHeaderToken)
	if !strings.Contains(headerDataToken, "Bearer") {
		err = constant.ErrHeaderTokenNotFound
		return
	}

	splitToken := strings.Split(headerDataToken, fmt.Sprintf("%s ", constant.DefaultMdwHeaderBearer))
	if len(splitToken) <= 1 {
		err = constant.ErrHeaderTokenInvalid
		return
	}

	token = splitToken[1]

	return
}
