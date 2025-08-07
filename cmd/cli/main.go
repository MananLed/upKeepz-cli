package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/MananLed/upKeepz-cli/internal/utils"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	userRepo := &repository.UserRepository{}
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	societyRepo := &repository.SocietyRepository{}
	societyService := service.NewSocietyService(societyRepo)
	societyHandler := handlers.NewSocietyHandler(societyService)

	credentialRepo := &repository.CredentialRepository{}
	credentialService := service.NewCredentialService(credentialRepo)
	credentialHandler := handlers.NewCredentialHandler(credentialService)

	noticeRepo := &repository.NoticeRepository{}
	noticeService := service.NewNoticeService(noticeRepo)
	noticeHandler := handlers.NewNoticeHandler(noticeService)

	serviceRequestRepo := &repository.ServiceRequestRepository{}
	serviceRequestService := service.NewServiceRequestService(serviceRequestRepo)
	serviceRequestHandler := handlers.NewServiceRequestHandler(serviceRequestService)

	feedbackRepo := &repository.FeedbackRepository{}
	feedbackService := service.NewFeedbackService(feedbackRepo)
	feedbackHandler := handlers.NewFeedbackHandler(feedbackService)

	invoiceRepo := &repository.InvoiceRepository{}
	invoiceService := service.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)

	reader := bufio.NewReader(os.Stdin)

	for {
		myFigure := figure.NewColorFigure("UpKeepz", "", "green", false)
		fmt.Println(constants.AppEmogiPrompt)
		myFigure.Print()
		fmt.Println(constants.AppEmogiPrompt)
		
		color.Cyan("1." + string(constants.SignUpPrompt))
		color.Cyan("2." + string(constants.LoginPrompt))
		color.Cyan("3." + string(constants.ExitPrompt))

		fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			userHandler.SignUp()
		case "2":
			user := userHandler.Login()
			if user == nil {
				continue
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, utils.UserIDKey, user.Email)
			ctx = context.WithValue(ctx, utils.UserRoleKey, user.Role)		
			ctx = context.WithValue(ctx, utils.UserPassKey, user.Password)
			switch user.Role {
			case model.RoleAdmin:
				ShowAdminDashboard(ctx, user, userHandler, societyHandler, credentialHandler, noticeHandler, feedbackHandler, invoiceHandler, serviceRequestHandler)
			case model.RoleOfficer:
				ShowOfficerDashboard(ctx, user, userHandler, serviceRequestHandler, noticeHandler, feedbackHandler, invoiceHandler)
			case model.RoleResident:
				ShowResidentDashboard(ctx, user, userHandler, serviceRequestHandler, noticeHandler, feedbackHandler, invoiceHandler)
			default:
				color.Green("Logged in as %s", user.Role)
			}
		case "3":
			color.Red("Exit")
			return
		default:
			color.Red("Invalid choice. Please try again.")
		}
	}
}