package service

import(
	"errors"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
)

type SocietyService struct{
	SocietyRepo repository.SocietyRepositoryInterface
}

func NewSocietyService(repo repository.SocietyRepositoryInterface) *SocietyService{
	return &SocietyService{SocietyRepo: repo}
}

func(s *SocietyService) GetAllResidents(requestSender model.User) ([]model.User, error){
	if requestSender.Role != model.RoleAdmin {
		return nil, errors.New("unauthorized access")
	}

	return s.SocietyRepo.GetAllResidents()
}

func (s *SocietyService) GetAllOfficers(requestSender model.User) ([]model.User, error){
	if requestSender.Role != model.RoleAdmin {
		return nil, errors.New("unauthorized access")
	}

	return s.SocietyRepo.GetAllOfficers()
}