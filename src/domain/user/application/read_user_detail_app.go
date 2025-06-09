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

func readUserDetailApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.ReadUserDetailRequest
		if err := c.ShouldBindUri(&request); err != nil {
			log.Printf("Error parsing URI param: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		log.Printf("Request to read user detail: %v", request.GUID)
		data, err := svc.ReadUserDetailService(c.Request.Context(), request)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedGetUserDetail,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToUserResponse(data),
			Message: msgSuccessGetUserDetail,
		})
	}
}
