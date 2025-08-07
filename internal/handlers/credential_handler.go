package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/fatih/color"
)

type CredentialHandler struct {
	Service *service.CredentialService
}

func NewCredentialHandler(s *service.CredentialService) *CredentialHandler {
	return &CredentialHandler{Service: s}
}

func (h *CredentialHandler) DeleteOfficer(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.YellowString("Enter the officer ID to delete: "))
	id, _ := reader.ReadString('\n')
	id = strings.TrimRight(id, "\r\n")
	err := h.Service.DeleteOfficerCredentials(ctx, id)
	if err != nil {
		color.Red("Failed to delete officer: %v", err)
	} else {
		color.Green("Officer credentials deleted successfully")
	}
}

func (h *CredentialHandler) DeleteResident(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.YellowString("Enter the resident ID to delete: "))

	id, _ := reader.ReadString('\n')
	id = strings.TrimRight(id, "\r\n")

	err := h.Service.DeleteResidentCredentials(ctx, id)

	if err != nil {
		color.Red("Failed to delete resident: %v", err)
	} else {
		color.Green("Resident credentials deleted successfully")
	}
}
