package service

import (
	"context"
	"errors"

	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/utils"
)

type CredentialService struct {
	Repo repository.CredentialRepositoryInterface
}

func NewCredentialService(r repository.CredentialRepositoryInterface) *CredentialService {
	return &CredentialService{Repo: r}
}

func (s *CredentialService) DeleteOfficerCredentials(ctx context.Context, officerID string) error {

	role, ok := ctx.Value(utils.UserRoleKey).(model.UserRole)

	if !ok || role != model.RoleAdmin {
		return errors.New("unauthorized: only admin can delete credentials")
	}

	return s.Repo.DeleteUserByIDAndRole(officerID, model.RoleOfficer)
}

func (s *CredentialService) DeleteResidentCredentials(ctx context.Context, residentID string) error {

	role, ok := ctx.Value(utils.UserRoleKey).(model.UserRole)

	if !ok || role != model.RoleAdmin {
		return errors.New("unauthorized: only admin can delete credentials")
	}
	
	return s.Repo.DeleteUserByIDAndRole(residentID, model.RoleResident)
}
