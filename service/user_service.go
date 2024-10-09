package service

import (
	"context"
	"go-booking/dto"
	"go-booking/entity"
	"go-booking/repository"
	"go-booking/utils"
	"sync"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) error
	}
	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

var (
	mu sync.Mutex
)

func (u *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) error {
	mu.Lock()
	defer mu.Unlock()

	_, err := u.userRepo.CheckEmail(ctx, req.Email)
	if err != nil {
		return dto.ErrEmailAlreadyExists
	}
	req.Password, _ = utils.HashPasword(req.Password)

	userData := entity.User{
		Name:       req.Name,
		TelpNumber: req.TelpNumber,
		Age:        req.Age,
		Email:      req.Email,
		Password:   req.Password,
	}

	_, err = u.userRepo.RegisterUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}
