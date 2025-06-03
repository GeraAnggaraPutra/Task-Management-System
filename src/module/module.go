package module

import (
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(
	db *gorm.DB,
) *Module {
	return &Module{
		db: db,
	}
}

func (k *Module) GetDB() *gorm.DB {
	return k.db
}
