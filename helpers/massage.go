package helpers

import "errors"

var (
	ErrIDNotFound           = errors.New("id not found")
	ErrEmailRequired        = errors.New("email is required")
	ErrEmailNotValid        = errors.New("email is not valid")
	ErrEmailHasBeenRegister = errors.New("email has been used")
	ErrPasswordRequired     = errors.New("password is required")

	CustomResponseMassageSuccess = "success"
	CustomResponseMassageFiled   = "Cannot acceess"
)
