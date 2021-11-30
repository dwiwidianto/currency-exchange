package response

import "currency-exchange/business/users"

type UsersCreateResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin int    `json:"is_admin" validate:"numberic"`
}

func FromDomainToCreateResponse(domain users.Domain) UsersCreateResponse {
	return UsersCreateResponse{
		Id:      domain.ID,
		Name:    domain.Name,
		Email:   domain.Email,
		IsAdmin: domain.IsAdmin,
	}
}
