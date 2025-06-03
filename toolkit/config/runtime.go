package config

import (
	"os"
	"time"

	"github.com/iancoleman/strcase"

	"task-management-system/src/util"
	"task-management-system/toolkit/constant"
)

type RuntimeConfig struct {
	Name                    string        `json:"name"`
	Host                    string        `json:"host"`
	Port                    int           `json:"port"`
	ShutdownTimeoutDuration time.Duration `json:"shutdown_timeout_duration"`
	ShutdownWaitDuration    time.Duration `json:"shutdown_wait_duration"`
	Mode                    string        `json:"mode"`
}

func NewRuntimeConfig() *RuntimeConfig {
	r := RuntimeConfig{}

	r.Name = os.Getenv("APP_NAME")
	r.Host = os.Getenv("APP_HOST")
	r.Port = util.ParseInt(constant.DefaultAppPort, os.Getenv("APP_PORT"))
	r.Mode = os.Getenv("APP_MODE")
	r.ShutdownTimeoutDuration = constant.DefaultAppShutdownTimeout
	r.ShutdownWaitDuration = constant.DefaultAppShutdownWait
	r.Name = strcase.ToSnake(r.Name)

	return &r
}
