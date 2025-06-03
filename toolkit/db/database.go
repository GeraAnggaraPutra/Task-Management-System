package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewDatabase() (db *gorm.DB, err error) {
	opt, err := newDatabaseOption()
	if err != nil {
		return
	}

	switch opt.driver {
	case "postgresql":
		db, err = NewPostgresql(opt)
	case "":
	default:
		err = errors.Wrapf(errors.New("invalid datasources driver"), "db: driver=%s", opt.driver)
	}

	return
}
