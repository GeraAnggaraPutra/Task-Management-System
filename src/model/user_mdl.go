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
	RoleGUID  string         `json:"role_guid"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
}

type UserDetail struct {
	GUID      string         `json:"guid" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	RoleGUID  string         `json:"role_guid"`
	RoleName  string         `json:"role_name"`
	Role      Role           `json:"roles" gorm:"foreignKey:RoleGUID;references:GUID"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
}

func (UserDetail) TableName() string {
	return "users"
}
