
// package main

// import(
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"github.com/MananLed/upKeepz-cli/internal/model"
// 	"github.com/MananLed/upKeepz-cli/internal/handlers"
// 	"github.com/fatih/color"
// 	"github.com/common-nighthawk/go-figure"
// 	"github.com/MananLed/upKeepz-cli/constants"
// )

// func ShowAdminDashboard(user *model.User, handler *handlers.SocietyHandler) {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println(constants.AdminEmogiPrompt)
// 		myFigure := figure.NewColorFigure("Admin","", "green", false)
// 		myFigure.Print()
// 		fmt.Println(constants.AdminEmogiPrompt)
// 		color.Cyan(string(constants.ViewResidentPrompt))
// 		color.Cyan(string(constants.ViewOfficersPrompt))
// 		color.Cyan(string(constants.LogoutPrompt))
// 		color.Blue(string(constants.ChoicePrompt))

// 		choice, _ := reader.ReadString('\n')
// 		choice = strings.TrimSpace(choice)

// 		switch choice {
// 		case "1":
// 			handler.HandleViewResidents(*user)
// 		case "2":
// 			handler.HandleViewOfficers(*user)
// 		case "3":
// 			color.Red("Logging out...")
// 			return
// 		default:
// 			color.Red("Invalid choice, try again.")
// 		}
// 	}
// }

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/fatih/color"
	"github.com/common-nighthawk/go-figure"
	"github.com/MananLed/upKeepz-cli/constants"
)

func ShowAdminDashboard(ctx context.Context, sHandler *handlers.SocietyHandler, cHandler *handlers.CredentialHandler, nHandler *handlers.NoticeHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(constants.AdminEmogiPrompt)
		myFigure := figure.NewColorFigure("Admin", "", "green", false)
		myFigure.Print()
		fmt.Println(constants.AdminEmogiPrompt)
		color.Cyan("1." + string(constants.ViewResidentPrompt))
		color.Cyan("2." + string(constants.ViewOfficersPrompt))
		color.Cyan("3." + string(constants.DeleteResidentPrompt))
		color.Cyan("4." + string(constants.DeleteOfficerPrompt))
		color.Cyan("5." + string(constants.IssueNoticePrompt))
		color.Cyan("6." + string(constants.GetNoticePrompt))
		color.Cyan("7." + string(constants.GetNoticeByID))
		color.Cyan("8." + string(constants.LogoutPrompt))
		color.Blue(string(constants.ChoicePrompt))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			sHandler.HandleViewResidents(ctx)
		case "2":
			sHandler.HandleViewOfficers(ctx)
		case "3":
			cHandler.DeleteResident(ctx)
		case "4":
			cHandler.DeleteOfficer(ctx)
		case "5":
			nHandler.IssueNotice(ctx)
		case "6":
			nHandler.GetNotices()
		case "7":
			nHandler.GetNoticeByID()
		case "8":
			color.Red("Logging out...")
			return
		default:
			color.Red("Invalid choice, try again.")
		}
	}
}

