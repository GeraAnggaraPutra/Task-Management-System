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

func readUserListApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.ReadUserListRequest
		if err := c.ShouldBind(&request); err != nil {
			log.Printf("Error parsing request: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		request.Init()

		data, totalData, err := svc.ReadUserListService(c.Request.Context(), request)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedGetUserList,
				Error:   &err,
			})
			return
		}

		module.ResponsePaginate(c, request.PaginationPayload, totalData, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToUserResponses(data),
			Message: msgSuccessGetUserList,
		})
	}
}
