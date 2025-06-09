package application

import (
	"github.com/gin-gonic/gin"

	"task-management-system/src/domain/user/service"
	"task-management-system/src/module"
	"task-management-system/src/middleware"

)

func AddRoutes(r *gin.Engine, m *module.Module) {
	svc := service.NewService(m.GetDB())
	mdw := middleware.NewEnsureToken(m.GetDB())

	routes := r.Group("/user", mdw.ValidateToken())

	routes.GET("", readUserListApp(svc))
	routes.GET("/:guid", readUserDetailApp(svc))
	routes.POST("", createUserApp(svc))
	routes.PUT("/:guid", updateUserApp(svc))
	routes.DELETE("/:guid", deleteUserApp(svc))
}
