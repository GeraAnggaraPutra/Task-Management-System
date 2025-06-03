package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"task-management-system/src/api"
	"task-management-system/src/module"
	"task-management-system/toolkit/config"
	"task-management-system/toolkit/db"
)

func main() {
	var err error

	// load .env file
	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %s", err.Error())
		}
	}

	ctx, cancel := config.NewRuntimeContext()
	defer func() {
		cancel()

		if err != nil {
			log.Printf("found error : %s", err.Error())
		}
	}()

	// setup database
	database, err := db.NewDatabase()
	if err != nil {
		return
	}

	// setup module
	mdl := module.NewModule(database)

	// run server
	api.RunGinServer(ctx, mdl)
}
