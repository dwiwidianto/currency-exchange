package users

import (
	"context"
	"currency-exchange/app/middlewares"
	"currency-exchange/helpers"
	"log"
	"time"
)

type UserUseCase struct {
	repo    UserRepoInterface
	ctx     time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration, jwtAuth *middlewares.ConfigJWT) *UserUseCase {
	return &UserUseCase{
		repo:    userRepo,
		ctx:     contextTimeout,
		jwtAuth: jwtAuth,
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
	data, err := usecase.repo.GetByEmailUsers(ctx, domain.Email)
	if data.ID > 0 {
		return Domain{}, helpers.ErrEmailHasBeenRegister
	}
	if domain.Password == "" {
		return Domain{}, helpers.ErrPasswordRequired
	}
	domain.Password, _ = helpers.HashPassword(domain.Password)
	user, err := usecase.repo.CreateUsers(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, string, error) {
	if email == "" || password == "" {
		return Domain{}, "", helpers.ErrUsernamePasswordNotFound
	}
	user, err := usecase.repo.GetByEmailUsers(ctx, email)

	if err != nil {
		return Domain{}, "", err
	}
	if !helpers.CheckPassword(password, user.Password) {
		return Domain{}, "", helpers.ErrInvalidAuthentication
	}

	token, errToken := usecase.jwtAuth.GenererateToken(user.ID, user.IsAdmin)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", helpers.ErrInvalidAuthentication
	}
	return user, token, nil
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
