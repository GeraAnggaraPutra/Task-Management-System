package application

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"task-management-system/src/constant"
	"task-management-system/src/domain/user/payload"
	"task-management-system/src/domain/user/service"
	"task-management-system/src/module"
)

func deleteUserApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.DeleteUserRequest
		if err := c.ShouldBindUri(&request); err != nil {
			log.Printf("Error parsing URI param: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		err := svc.DeleteUserService(c.Request.Context(), request)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedDeleteUser,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteUser,
		})
	}
}
