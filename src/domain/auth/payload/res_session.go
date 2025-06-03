package payload

import (
	"time"

	"task-management-system/src/model"
)

type UserResponse struct {
	GUID      string  `json:"guid"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	CreatedBy *string `json:"created_by"`
	UpdatedAt *string `json:"updated_at"`
	UpdatedBy *string `json:"updated_by"`
}

type SessionResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  UserResponse `json:"user"`
}

func ToSessionResponse(entity model.Session, user model.User) (response SessionResponse) {
	response.AccessToken = entity.AccessToken
	response.AccessTokenExpiresAt = entity.AccessTokenExpiresAt
	response.RefreshToken = entity.RefreshToken
	response.RefreshTokenExpiresAt = entity.RefreshTokenExpiresAt
	response.User.GUID = user.GUID
	response.User.Username = user.Username
	response.User.Email = user.Email
	response.User.CreatedAt = user.CreatedAt.Format(time.RFC3339)

	if user.CreatedBy.Valid {
		response.User.CreatedBy = &user.CreatedBy.String

	}

	if user.UpdatedAt.Valid {
		updatedAt := user.UpdatedAt.Time.Format(time.RFC3339)
		response.User.UpdatedAt = &updatedAt
	}

	if user.UpdatedBy.Valid {
		response.User.UpdatedBy = &user.UpdatedBy.String
	}

	return
}
