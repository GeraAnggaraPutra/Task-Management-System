package payload

type ReadUserDetailRequest struct {
	GUID string `uri:"guid" validate:"required"`
}
