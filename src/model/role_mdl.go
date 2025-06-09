package model

import (
	"database/sql"
	"time"
)

type Role struct {
	GUID        string         `json:"guid" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   sql.NullString `json:"created_by"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	UpdatedBy   sql.NullString `json:"updated_by"`
}
