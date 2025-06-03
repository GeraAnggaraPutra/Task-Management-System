package application

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"task-management-system/src/constant"
	"task-management-system/src/domain/auth/payload"
	"task-management-system/src/domain/auth/service"
	"task-management-system/src/module"
)

func refreshTokenApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.RefreshTokenRequest
		if err := c.ShouldBind(&request); err != nil {
			log.Printf("Error parsing request: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		data, user, err := svc.RefreshTokenService(c.Request.Context(), request)
		if err != nil {
			log.Printf("Error refreshing token: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusInternalServerError,
				Message: msgFailedRefreshToken,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(data, user),
			Message: msgSuccessRefreshToken,
		})
	}
}
