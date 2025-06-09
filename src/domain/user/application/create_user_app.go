package application

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"task-management-system/src/constant"
	"task-management-system/src/domain/user/payload"
	"task-management-system/src/domain/user/service"
	"task-management-system/src/handler/auth"
	"task-management-system/src/module"
)

func createUserApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payload.CreateUserRequest
		if err := c.ShouldBind(&request); err != nil {
			log.Printf("Error parsing request: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: constant.ErrFailedParseRequest.Error(),
				Error:   &err,
			})
			return
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Error:   &err,
			})
			return
		}

		exists, err := svc.CheckEmailIsExistsService(c.Request.Context(), request.Email)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusInternalServerError,
				Message: msgFailedCheckEmailExists,
				Error:   &err,
			})
			return
		}

		if exists {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgEmailAlreadyExists,
				Error:   nil,
			})
			return
		}

		err = svc.CreateUserService(c.Request.Context(), request, ah.GetClaims().UserGUID)
		if err != nil {
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedCreateUser,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessCreateUser,
		})
	}
}
