package constant

import (
	"errors"
)

// error definitions
var (
	ErrFailedParseRequest = errors.New("failed to parse request")

	ErrHeaderTokenNotFound = errors.New("header authorization not found")
	ErrHeaderTokenInvalid  = errors.New("invalid header token")
	ErrTokenInvalid        = errors.New("invalid token")
	ErrTokenMissing        = errors.New("missing token")
	ErrTokenExpired        = errors.New("expired token")
	ErrTokenUnauthorized   = errors.New("unauthorized token")
	ErrUserNotFound        = errors.New("user not found")

	ErrForbiddenPermission = errors.New("your permission is not allowed to access this resource")
	ErrForbiddenRole       = errors.New("your role is not allowed to access this resource")

	ErrDataNotFound = errors.New("data not found")

	ErrUnknownSource = errors.New("an error occurred, please try again later")
)

// error message.
const (
	ErrMsgValidate      = "There are some errors in your request"
	ErrMsgUnknownSource = "an error occurred, please try again later"
)

// error form field.
var (
	// password.
	ErrPasswordIncorrect = errors.New("password incorrect")

	// email.
	ErrAccountNotFound        = errors.New("account not found")
	ErrAccountNotHavePassword = errors.New("account does not have password")
	ErrEmailAlreadyExists     = errors.New("this email has been used")
)
