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

func loginApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.LoginRequest
		if err := c.ShouldBind(&request); err != nil {
			log.Printf("Error parsing request: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		data, user, err := svc.LoginService(c.Request.Context(), request)
		if err != nil {
			log.Printf("Error login: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedLogin,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(data, user),
			Message: msgSuccessLogin,
		})
	}
}
