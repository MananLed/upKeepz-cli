package repository

import(
	"encoding/json"
	"errors"
	"os"
	"sync"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

type UserRepositoryInterface interface {
	AddUser(user model.User) error
	GetUserByID(id string) (*model.User, error)
}

type UserRepository struct{
	mu sync.Mutex 
}

func (r *UserRepository) LoadUsers() ([]model.User, error){
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(userDataFile)

	if err != nil{
		if os.IsNotExist(err){
			return []model.User{}, nil
		}
		return nil, err 
	}
	
	var users []model.User
	err = json.Unmarshal(data, &users)
	if err != nil{
		return nil, err 
	}

	return users, nil 
}

func (r *UserRepository) SaveUsers(users []model.User) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := json.MarshalIndent(users, ""," ")
	if err != nil{
		return err
	}
	return os.WriteFile(userDataFile, data, 0644)
}

func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	users, err := r.LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) AddUser(newUser model.User) error {
	users, err := r.LoadUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.ID == newUser.ID {
			return errors.New("user with this ID already exists")
		}
	}

	users = append(users, newUser)
	return r.SaveUsers(users)
}