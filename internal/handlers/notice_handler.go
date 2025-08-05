package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/MananLed/upKeepz-cli/internal/utils"
	"github.com/fatih/color"
)

type NoticeHandler struct {
	NoticeService *service.NoticeService
}

func NewNoticeHandler(service *service.NoticeService) *NoticeHandler {
	return &NoticeHandler{NoticeService: service}
}

func (h *NoticeHandler) IssueNotice(ctx context.Context) {
	user , err := utils.GetUserFromContext(ctx)

	if user.Role == model.RoleResident{
		color.Red("not permitted to issue notice.")
		return
	}
	if err != nil {
		color.Red("error:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Enter notice content:")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	if content == "" {
		color.Red("Notice content cannot be empty.")
		return
	}

	err = h.NoticeService.IssueNotice(content)
	if err != nil {
		color.Red("Failed to issue notice: %v", err)
	} else {
		color.Green("Notice issued successfully.")
	}
}

func (h *NoticeHandler) GetNotices() {
	notices, err := h.NoticeService.GetNotices()
	if err != nil {
		color.Red("Failed to retrieve notices: %v", err)
		return
	}

	if len(notices) == 0 {
		color.Yellow("No notices found.")
		return
	}

	color.Cyan("===== All Notices =====")
	for _, n := range notices {
		fmt.Printf("ID: %d\t|\tDate: %s\nContent: \n %s\n\n", n.ID, n.DateIssued, n.Content)
	}
}
func (h *NoticeHandler) GetNoticeByID() {
	color.Yellow("Enter notice ID:")
	var id int
	_, err := fmt.Scanf("%d\n", &id)

	if err != nil {
		color.Red("Invalid input. Please enter a number.")
		return
	}

	notice, err := h.NoticeService.GetNoticeByID(id)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	color.Green("Notice Found:")
	fmt.Printf("ID: %d\t|\tDate: %s\nContent: \n %s\n\n", notice.ID, notice.DateIssued, notice.Content)
}

