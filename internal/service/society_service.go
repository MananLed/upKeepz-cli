package service

import (
	"context"
	"errors"

	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/utils"
)

type SocietyService struct {
	SocietyRepo repository.SocietyRepositoryInterface
}

func NewSocietyService(repo repository.SocietyRepositoryInterface) *SocietyService {
	return &SocietyService{SocietyRepo: repo}
}

func (s *SocietyService) GetAllResidents(ctx context.Context) ([]model.User, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role != model.RoleAdmin {
		return nil, errors.New("unauthorized access: only admins can view residents")
	}
	return s.SocietyRepo.GetAllResidents()
}

func (s *SocietyService) GetAllOfficers(ctx context.Context) ([]model.User, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role != model.RoleAdmin {
		return nil, errors.New("unauthorized access: only admins can view officers")
	}
	return s.SocietyRepo.GetAllOfficers()
}
