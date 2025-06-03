package application

import (
	"github.com/gin-gonic/gin"

	"task-management-system/src/domain/auth/service"
	"task-management-system/src/middleware"
	"task-management-system/src/module"
)

func AddRoutes(r *gin.Engine, m *module.Module) {
	svc := service.NewService(m.GetDB())
	mdw := middleware.NewEnsureToken(m.GetDB())

	routes := r.Group("/auth")

	routes.POST("/login", loginApp(svc))
	routes.POST("/refresh-token", refreshTokenApp(svc))
	routes.POST("/logout", mdw.ValidateToken(), logoutApp(svc))
}
