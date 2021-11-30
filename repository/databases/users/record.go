package users

import (
	"currency-exchange/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Email     string
	Name      string
	Password  string
	IsAdmin   int            `gorm:"default:0"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (rec *User) UserToDomain() (res *users.Domain) {
	if rec != nil {
		res = &users.Domain{
			ID:        rec.ID,
			Name:      rec.Name,
			Email:     rec.Email,
			Password:  rec.Password,
			IsAdmin:   rec.IsAdmin,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res

}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		IsAdmin:   domain.IsAdmin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func UsersToListDomain(data []User) []users.Domain {
	var list []users.Domain
	for _, v := range data {
		list = append(list, *v.UserToDomain())
	}
	return list
}
