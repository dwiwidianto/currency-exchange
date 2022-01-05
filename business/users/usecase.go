package users

import (
	"context"
	"currency-exchange/helpers"
	"errors"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
}

func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration) *UserUseCase {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *UserUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := usecase.repo.GetAllUsers(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) Create(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, helpers.ErrEmailRequired
	}
	if !helpers.IsEmailValid(domain.Email) {
		return Domain{}, helpers.ErrEmailNotValid
	}
	data, _ := usecase.repo.GetByEmailUsers(ctx, domain.Email)

	if data.ID > 0 {
		return Domain{}, helpers.ErrEmailHasBeenRegister
	}
	if domain.Password == "" {
		return Domain{}, helpers.ErrPasswordRequired
	}
	// domain.Password, _ = helpers.HashPassword(domain.Password)
	user, err := usecase.repo.CreateUsers(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := usecase.repo.Login(domain, ctx)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) GetById(ctx context.Context, id int) (Domain, error) {
	user, err := usecase.repo.GetByIdUsers(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.ID == 0 {
		return Domain{}, helpers.ErrIDNotFound
	}
	return user, nil
}

func (usecase *UserUseCase) Delete(ctx context.Context, id int) error {
	err := usecase.repo.DeleteUsers(ctx, id)
	if err != nil {
		return err
	}
	return nil

}
