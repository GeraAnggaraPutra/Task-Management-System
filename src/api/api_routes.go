package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	authDomain "task-management-system/src/domain/auth/application"
	userDomain "task-management-system/src/domain/user/application"
	"task-management-system/src/module"
)

func routes(r *gin.Engine, m *module.Module) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": os.Getenv("APP_NAME") + " is Running",
		})
	})

	// domain routes
	authDomain.AddRoutes(r, m)
	userDomain.AddRoutes(r, m)
}
