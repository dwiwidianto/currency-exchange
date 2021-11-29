package users

import (
	"context"
	"currency-exchange/business/users"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewPosgresUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(&domain)
	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, errors.New("Email not found")
		}
		return users.Domain{}, errors.New("Error in database")
	}
	return *userDb.UserToDomain(), nil
}

func (repo *UserRepository) FindByID(ctx context.Context, userID int) (users.Domain, error) {
	rec := User{}
	err := repo.db.Where("user.id = ?", userID).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return *rec.UserToDomain(), nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	var data []User
	err := repo.db.Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return UsersToListDomain(data), nil
}

func (repo *UserRepository) GetByIdUsers(ctx context.Context, id int) (users.Domain, error) {
	var user User
	err := repo.db.Find(&user, "user.id = ?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return *user.UserToDomain(), nil
}

func (repo *UserRepository) GetByEmailUsers(ctx context.Context, email string) (users.Domain, error) {
	var data User
	err := repo.db.Find(&data, "email=?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return *data.UserToDomain(), nil
}

func (repo *UserRepository) CreateUsers(ctx context.Context, register *users.Domain) (users.Domain, error) {
	user := User{
		Name:     register.Name,
		Email:    register.Email,
		Password: register.Password,
	}
	err := repo.db.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return *user.UserToDomain(), nil
}

func (repo *UserRepository) DeleteUsers(ctx context.Context, id int) error {
	user := User{}
	err := repo.db.Delete(&user, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
