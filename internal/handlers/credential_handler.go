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

func promptInput(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (h *CredentialHandler) DeleteOfficer(ctx context.Context) {
	id := promptInput("Enter Officer ID to delete: ")
	err := h.Service.DeleteOfficerCredentials(ctx, id)
	if err != nil {
		color.Red("Failed to delete officer: %v", err)
	} else {
		color.Green("Officer credentials deleted successfully")
	}
}

func (h *CredentialHandler) DeleteResident(ctx context.Context) {
	id := promptInput("Enter Resident ID to delete: ")
	err := h.Service.DeleteResidentCredentials(ctx, id)
	if err != nil {
		color.Red("Failed to delete resident: %v", err)
	} else {
		color.Green("Resident credentials deleted successfully")
	}
}
