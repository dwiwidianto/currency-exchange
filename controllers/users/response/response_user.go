package response

import (
	"currency-exchange/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromUsersDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromUserListDomain(domain []users.Domain) []UserResponse {
	var list []UserResponse
	for _, v := range domain {
		list = append(list, FromUsersDomain(v))
	}
	return list
}
