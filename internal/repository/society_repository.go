package repository

import (
	"encoding/json"
	"os"
	"fmt"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

type SocietyRepository struct{}

type SocietyRepositoryInterface interface{
	GetAllResidents() ([]model.User, error)
	GetAllOfficers() ([]model.User, error)
}

func (s *SocietyRepository) GetAllUsers() ([]model.User, error){
	
	data, err := os.ReadFile(userDataFile)

	if err != nil{
		return nil, err 
	}
    
	var users []model.User

	err = json.Unmarshal(data, &users)

	if err != nil{
		return nil, err
	}

	return users, nil 
}

func (s *SocietyRepository) GetAllResidents() ([]model.User, error){
	users, err := s.GetAllUsers()

	if err != nil{
		return nil, err
	}

	var residents []model.User 
	var count int = 0
	for _, user := range users{
		if user.Role == model.RoleResident {
			count++
			residents = append(residents, user)
		}
	}
	fmt.Println("Total Residents:", count)
	return residents, nil 
}

func (s *SocietyRepository) GetAllOfficers() ([]model.User, error){
	users, err := s.GetAllUsers()
	
	if err != nil {
		return nil, err 
	}

	var officers []model.User 
	var count int = 0
	for _, user := range users{
		if user.Role == model.RoleOfficer {
			count++
			officers = append(officers, user)
		}
	}
	fmt.Println("Total Officers:", count)
	return officers, nil
}



