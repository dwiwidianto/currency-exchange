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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (rec *User) UserToDomain() (res *users.Domain) {
	if rec != nil {
		res = &users.Domain{
			ID:        rec.ID,
			Name:      rec.Name,
			Email:     rec.Email,
			Password:  rec.Password,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res

}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func UsersToListDomain(data []User) []users.Domain {
	var list []users.Domain
	for _, v := range data {
		list = append(list, *v.UserToDomain())
	}
	return list
}
