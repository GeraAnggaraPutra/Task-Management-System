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

	if cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// Register Routes
	routes(r, m)

	log.Printf("serving REST HTTP server : %s", logger.ParseJSON(cfg))

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ERROR starting HTTP server: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down the server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeoutDuration)
	defer cancel()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("ERROR shutting down server: %v", err)
	}

	log.Println("Server exited gracefully")
}
