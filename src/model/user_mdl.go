package model

import (
	"database/sql"
	"time"
)

type User struct {
	GUID      string         `json:"guid" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
}
