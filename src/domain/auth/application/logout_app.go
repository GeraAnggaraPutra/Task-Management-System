package application

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"task-management-system/src/domain/auth/service"
	"task-management-system/src/handler/auth"
	"task-management-system/src/module"
)

func logoutApp(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ah, err := auth.GetAuth(c)
		if err != nil {
			return
		}

		err = svc.LogoutService(c.Request.Context(), ah.GetClaims())
		if err != nil {
			log.Printf("Error logout: %v", err)
			module.ResponseData(c, module.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: msgFailedLogout,
				Error:   &err,
			})
			return
		}

		module.ResponseData(c, module.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessLogout,
		})
	}
}
