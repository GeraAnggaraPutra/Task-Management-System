package payload

import (
	"time"

	"task-management-system/src/model"
)

type UserResponse struct {
	GUID      string     `json:"guid"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	RoleGUID  string     `json:"role_guid"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

func ToUserResponse(entity model.UserDetail) (response UserResponse) {
	response.GUID = entity.GUID
	response.Username = entity.Username
	response.Email = entity.Email
	response.RoleGUID = entity.RoleGUID
	response.RoleName = entity.Role.Name
	response.CreatedAt = entity.CreatedAt

	if entity.Role.Name == "" {
		response.RoleName = entity.RoleName
	}

	if entity.CreatedBy.Valid {
		response.CreatedBy = &entity.CreatedBy.String
	}

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

func ToUserResponses(entities []model.UserDetail) (response []UserResponse) {
	response = make([]UserResponse, len(entities))

	for i := range entities {
		response[i] = ToUserResponse(entities[i])
	}

	return
}
