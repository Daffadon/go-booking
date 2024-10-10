package service

import (
	"context"
	"fmt"
	"go-booking/dto"
	"go-booking/entity"
	"go-booking/repository"
	"go-booking/utils"
	"sync"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) error
		LoginUser(ctx context.Context, req dto.UserLoginRequest) (string, error)
	}
	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

var (
	mu sync.Mutex
)

func (u *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) error {
	mu.Lock()
	defer mu.Unlock()

	_, err := u.userRepo.CheckIsEmailExist(ctx, req.Email)
	if err == nil {
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

func (u *userService) LoginUser(ctx context.Context, req dto.UserLoginRequest) (string, error) {
	user, err := u.userRepo.CheckIsEmailExist(ctx, req.Email)
	if err != nil {
		return "", dto.ErrEmailOrPasswordIsWrong
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		fmt.Print(err)
		return "", dto.ErrEmailOrPasswordIsWrong
	}
	token := u.jwtService.GenerateToken(user.ID.String(), user.Role)
	return token, nil
}
