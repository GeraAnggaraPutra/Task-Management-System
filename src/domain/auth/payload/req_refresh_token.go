package payload

import "task-management-system/src/model"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (request *RefreshTokenRequest) ToSessionPayload(session model.Session) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: session.GUID,
		UserGUID:    session.UserGUID,
	}

	return
}
