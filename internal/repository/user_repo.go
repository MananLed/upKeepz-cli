package repository

import(
	"encoding/json"
	"errors"
	"os"
	"sync"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

const dataDir = "../../data"
const userDataFile = dataDir + "/users.json"

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

	file, err := os.Open(userDataFile)

	if err != nil{
		if os.IsNotExist(err){
			return []model.User{}, nil
		}
		return nil, err 
	}
	defer file.Close()

	var users []model.User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil{
		return nil, err 
	}

	return users, nil 
}

func (r *UserRepository) SaveUsers(users []model.User) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	_ = os.MkdirAll(dataDir, os.ModePerm)

	file, err := os.Create(userDataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	return encoder.Encode(users)
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