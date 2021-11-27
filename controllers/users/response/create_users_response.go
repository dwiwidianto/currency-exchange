package response

import "currency-exchange/business/users"

type UsersCreateResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomainToCreateResponse(domain users.Domain) UsersCreateResponse {
	return UsersCreateResponse{
		Id:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
	}
}
