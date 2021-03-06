package request

import (
	"currency-exchange/business/users"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (ur *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     ur.Name,
		Email:    ur.Email,
		Password: ur.Password,
	}
}
