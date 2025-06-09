package payload

type UserPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	RoleGUID string `json:"role_guid" validate:"required"`
}
