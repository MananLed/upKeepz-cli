package handlers

import (
	"context"
	"fmt"

	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/MananLed/upKeepz-cli/internal/utils"
)

type SocietyHandler struct {
	SocietyService *service.SocietyService
}

func NewSocietyHandler(serve *service.SocietyService) *SocietyHandler {
	return &SocietyHandler{SocietyService: serve}
}

func (h *SocietyHandler) HandleViewResidents(ctx context.Context) {
	_, err := utils.GetUserFromContext(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	residents, err := h.SocietyService.GetAllResidents(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if (len(residents) > 0) {fmt.Println("All Residents:-")}
	for _, user := range residents {
		fmt.Printf("ID: %s\n", user.ID)
	}
}

func (h *SocietyHandler) HandleViewOfficers(ctx context.Context) {
	_, err := utils.GetUserFromContext(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	officers, err := h.SocietyService.GetAllOfficers(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if(len(officers) > 0) {fmt.Println("All Officers:-")}
	for _, user := range officers {
		fmt.Printf("ID: %s\n", user.ID)
	}
}
