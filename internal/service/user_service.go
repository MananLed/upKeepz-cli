package service

import(
	"errors"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
)

type UserService struct{
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService{
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

func (us *UserService) Login(id string, password string) (*model.User, error){
	user, err := us.UserRepo.GetUserByID(id)

	if err != nil{
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil 
}