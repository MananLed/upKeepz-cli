package repository

import (
	"errors"
	"fmt"

	"github.com/MananLed/upKeepz-cli/internal/model"
)

type CredentialRepository struct {
	UserRepository
}

type CredentialRepositoryInterface interface {
	DeleteUserByIDAndRole(id string, role model.UserRole) error
}

func (r *CredentialRepository) DeleteUserByIDAndRole(id string, role model.UserRole) error {
	users, err := r.LoadUsers()
	if err != nil {
		return err
	}

	var updated []model.User
	found := false

	fmt.Println(id, role)

	for _, u := range users {
		fmt.Println(u.ID, u.Role)
		if u.ID == id && u.Role == role {
			found = true
			continue
		}
		updated = append(updated, u)
	}

	if !found {
		return errors.New("user not found")
	}

	return r.SaveUsers(updated)
}
