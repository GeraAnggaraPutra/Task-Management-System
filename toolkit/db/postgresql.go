package db

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresql(opt *databaseOption) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", opt.host, opt.port, opt.username, opt.password, opt.schema, opt.sslmode)

	var cfgLogger logger.Interface

	if opt.isLog {
		cfgLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             opt.slowThreshold,
				LogLevel:                  opt.level,
				IgnoreRecordNotFoundError: opt.ignoreErr,
				Colorful:                  opt.colorful,
			},
		)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: cfgLogger})
	if err != nil {
		err = errors.Wrap(err, "gorm: failed to open connection")
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		err = errors.Wrap(err, "gorm: failed to get generic database object")
		return
	}

	sqlDB.SetMaxIdleConns(opt.connectionOption.maxIdle)
	sqlDB.SetMaxOpenConns(opt.connectionOption.maxOpen)
	sqlDB.SetConnMaxLifetime(opt.connectionOption.maxLifetime)

	log.Printf("successfully connected to postgresql %s:%d", opt.host, opt.port)

	go keepAlive(sqlDB, opt.driver, opt.schema, opt.keepAliveInterval)

	return
}
