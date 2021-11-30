package request

import (
	"currency-exchange/business/users"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(login UserLogin) users.Domain {
	return users.Domain{
		Email:    login.Email,
		Password: login.Password,
	}
}
