package payload

type DeleteUserRequest struct {
	GUID string `uri:"guid" validate:"required"`
}
