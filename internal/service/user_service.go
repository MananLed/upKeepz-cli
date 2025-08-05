package service

import (
	"context"
	"errors"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/utils"
)

type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (us *UserService) SignUp(user model.User) error {
	if strings.TrimSpace(user.ID) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("ID and Password cannot be empty")
	}
	return us.UserRepo.AddUser(user)
}

func (us *UserService) Login(id string, password string) (*model.User, error) {
	user, err := us.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user.Password != password {
		return nil, errors.New("incorrect password")
	}
	return user, nil
}

// Utility function if you want to load user from context in future services
func (us *UserService) GetLoggedInUser(ctx context.Context) (*model.User, error) {
	return utils.GetUserFromContext(ctx)
}
