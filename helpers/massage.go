package helpers

import "errors"

var (
	ErrIDNotFound               = errors.New("id not found")
	ErrUsernamePasswordNotFound = errors.New("username or password empty")
	ErrEmailRequired            = errors.New("email is required")
	ErrEmailNotValid            = errors.New("email is not valid")
	ErrEmailHasBeenRegister     = errors.New("email has been used")
	ErrPasswordRequired         = errors.New("password is required")
	ErrInvalidAuthentication    = errors.New("authentication failed: invalid user credentials")
	ErrInvalidTokenCredential   = errors.New("token not found or expired")

	CustomResponseMassageSuccess = "success"
	CustomResponseMassageFiled   = "Cannot acceess"
)
