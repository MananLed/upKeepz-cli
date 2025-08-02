package service

import(
	"errors"
	"testing"
	"github.com/MananLed/upKeepz-cli/internal/model" 
)

type MockUserRepo struct{
	users map[string]model.User
}

func (m *MockUserRepo) AddUser(user model.User) error {
	if _, exists := m.users[user.ID]; exists{
		return errors.New("user already exists")
	}
	m.users[user.ID] = user 
	return nil 
}

func (m *MockUserRepo) GetUserByID(id string) (*model.User, error) {
	user, exists := m.users[id] 
	if !exists{
		return nil, errors.New("user not found")
	}
	return &user, nil 
}

//Tests

func TestSignUp(t *testing.T){
	mockRepo := &MockUserRepo{users: make(map[string]model.User)}
	service := NewUserService(mockRepo)

	user := model.User{
		ID: "man",
		Password: "nin",
		Role: model.RoleResident,
	}

	err := service.SignUp(user)

	if err != nil{
		t.Errorf("expected signup to succeed, got error: %v", err)
	}
}

func TestLogin(t *testing.T){
	mockRepo := &MockUserRepo{users: make(map[string]model.User)}
	service := NewUserService(mockRepo)

	user := model.User{
		ID: "man",
		Password: "nin",
	}

	_ = service.SignUp(user)

	loggedInUser, err := service.Login("man", "nin")

	if err != nil || loggedInUser.ID != "man" {
		t.Errorf("login failed: %v", err)
	}
}

func TestLoginFailWrongPassword(t *testing.T){
	mockRepo := &MockUserRepo{users: make(map[string]model.User)}
	service := NewUserService(mockRepo)

	user := model.User{ID: "man", Password: "nin"}

	_ = service.SignUp(user)

	_, err := service.Login("man", "sin")

	if err == nil{
		t.Error("expected login to fail due to wrong password")
	}
}