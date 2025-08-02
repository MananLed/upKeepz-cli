package service

import(
	"testing"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

type MockSocietyRepo struct{
	Users []model.User 
}

func (m *MockSocietyRepo) GetAllResidents() ([]model.User, error){
	var residents []model.User
	for _, u := range m.Users {
		if u.Role == model.RoleResident{
			residents = append(residents, u)
		}
	}
	return residents, nil
}

func (m *MockSocietyRepo) GetAllOfficers() ([]model.User, error){
	var officers []model.User
	for _, u := range m.Users {
		if u.Role == model.RoleOfficer{
			officers = append(officers, u)
		}
	}
	return officers, nil
}

func TestGetAllResidentsAsAdmin(t *testing.T){
	mockRepo := &MockSocietyRepo{
		Users : []model.User{
			{
				FirstName:    "Meera",
    			MiddleName:   "Rani",
				LastName:     "Verma",
				MobileNumber: "9123456789",
				Email:        "meera.verma@example.com",
				ID:           "r1",
				Password:     "Pass#456",
				Role:         model.RoleResident,
			},	
			{
				
				FirstName:    "Rohan",
				MiddleName:   "Dev",
				LastName:     "Patel",
				MobileNumber: "9988776655",
				Email:        "rohan.patel@example.com",
				ID:           "o1",
				Password:     "MyPass789!",
				Role:         model.RoleOfficer,
			},
		},
	}

	service := NewSocietyService(mockRepo)

	admin := model.User{
		
		FirstName:    "Aarav",
    	MiddleName:   "Kumar",
		LastName:     "Sharma",
		MobileNumber: "9876543210",
		Email:        "aarav.sharma@example.com",
		ID:           "a1",
		Password:     "Secure@123",
		Role:         model.RoleAdmin, 
	}

	residents, err := service.GetAllResidents(admin)

	if err != nil {
		t.Fatalf("expected residents list but got error %v", residents)
	}
	if len(residents) != 1 || residents[0].ID != "r1" {
		t.Errorf("unexpected residents list: %v", residents)
	}
}

func TestGetAllOfficersAsAdmin(t *testing.T){
	mockRepo := &MockSocietyRepo{
		Users : []model.User{
			{
				FirstName:    "Meera",
    			MiddleName:   "Rani",
				LastName:     "Verma",
				MobileNumber: "9123456789",
				Email:        "meera.verma@example.com",
				ID:           "r1",
				Password:     "Pass#456",
				Role:         model.RoleResident,
			},	
			{
				FirstName:    "Rohan",
				MiddleName:   "Dev",
				LastName:     "Patel",
				MobileNumber: "9988776655",
				Email:        "rohan.patel@example.com",
				ID:           "o1",
				Password:     "MyPass789!",
				Role:         model.RoleOfficer,
			},
		},
	}

	service := NewSocietyService(mockRepo)

	admin := model.User{
		FirstName:    "Aarav",
    	MiddleName:   "Kumar",
		LastName:     "Sharma",
		MobileNumber: "9876543210",
		Email:        "aarav.sharma@example.com",
		ID:           "a1",
		Password:     "Secure@123",
		Role:         model.RoleAdmin, 
	}

	officers, err := service.GetAllOfficers(admin)

	if err != nil {
		t.Fatalf("expected officers list but got error %v", officers)
	}
	if len(officers) != 1 || officers[0].ID != "o1" {
		t.Errorf("unexpected residents list: %v", officers)
	}
}

func TestGetAllResidentsAsNonAdmin(t *testing.T){
	mockRepo := &MockSocietyRepo{}
	service := NewSocietyService(mockRepo)

	officer := model.User{
				FirstName:    "Rohan",
				MiddleName:   "Dev",
				LastName:     "Patel",
				MobileNumber: "9988776655",
				Email:        "rohan.patel@example.com",
				ID:           "o1",
				Password:     "MyPass789!",
				Role:         model.RoleOfficer,
			}

	_, err := service.GetAllResidents(officer)
	if err == nil {
		t.Error("expected error for non-admin user")
	}
}

func TestGettAllOfficersAsNonAdmin(t *testing.T){
	mockRepo := &MockSocietyRepo{}
	service := NewSocietyService(mockRepo)

	resident := model.User{
				FirstName:    "Meera",
    			MiddleName:   "Rani",
				LastName:     "Verma",
				MobileNumber: "9123456789",
				Email:        "meera.verma@example.com",
				ID:           "r1",
				Password:     "Pass#456",
				Role:         model.RoleResident,
	}
	
	_, err := service.GetAllOfficers(resident)
	if err == nil {
		t.Error("expected error for non-admin user")
	}
}