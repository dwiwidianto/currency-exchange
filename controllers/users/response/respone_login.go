package response

import (
	"currency-exchange/business/users"
)

func FromUserToDomainLogin(domain users.Domain, token string) LoginResponse {
	response := UserResponse{
		ID:      domain.ID,
		Name:    domain.Name,
		Email:   domain.Email,
		IsAdmin: domain.IsAdmin,
	}
	loginResponse := LoginResponse{}
	loginResponse.SessionToken = token
	loginResponse.User = response
	return loginResponse
}
