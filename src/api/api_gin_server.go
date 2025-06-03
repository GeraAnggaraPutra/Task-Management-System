package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"task-management-system/src/module"
	"task-management-system/toolkit/config"
	"task-management-system/toolkit/logger"
)

func RunGinServer(ctx context.Context, m *module.Module) {
	cfg := config.NewRuntimeConfig()

	r := gin.Default()

	// Register Routes
	routes(r, m)

	log.Printf("serving REST HTTP server : %s", logger.ParseJSON(cfg))

	if err := r.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ERROR starting HTTP server: %v", err)
	}

	log.Println("Server stopped.")
}
