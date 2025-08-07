package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/fatih/color"
)

type SocietyRepository struct{}

type SocietyRepositoryInterface interface{
	GetAllResidents() ([]model.User, error)
	GetAllOfficers() ([]model.User, error)
}

func (s *SocietyRepository) GetAllUsers() ([]model.User, error){
	
	data, err := os.ReadFile(string(constants.UserDataPath))

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
	if(len(residents) != 0) {
		fmt.Print(color.YellowString("Total Residents: "), count)
	}else {
		fmt.Print("There are no residents currently.")
	}
	fmt.Println()
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
	if(len(officers) != 0) {
		fmt.Print(color.YellowString("Total Officers: "), count)
	} else{
		fmt.Print("There are no officers currently.")
	}
	fmt.Println()
	return officers, nil
}