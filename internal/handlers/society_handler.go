package handlers

import(
	"fmt"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
)

type SocietyHandler struct{
	SocietyService *service.SocietyService 
}

func NewSocietyHandler(serve *service.SocietyService) *SocietyHandler {
	return &SocietyHandler{SocietyService: serve}
}

func (h *SocietyHandler) HandleViewResidents(currentUser model.User){
	residents, err := h.SocietyService.GetAllResidents(currentUser)

	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All Residents:")
	for _, user := range residents {
		fmt.Printf("ID: %s\n", user.ID)
	}
}

func (h *SocietyHandler) HandleViewOfficers(currentUser model.User) {
	officers, err := h.SocietyService.GetAllOfficers(currentUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All Officers:")
	for _, user := range officers {
		fmt.Printf("ID: %s\n", user.ID)
	}
}