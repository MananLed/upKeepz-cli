
package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/fatih/color"
	"github.com/common-nighthawk/go-figure"
	"github.com/MananLed/upKeepz-cli/constants"
)

func ShowAdminDashboard(user *model.User, handler *handlers.SocietyHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(constants.AdminEmogiPrompt)
		myFigure := figure.NewColorFigure("Admin","", "green", false)
		myFigure.Print()
		fmt.Println(constants.AdminEmogiPrompt)
		color.Cyan(string(constants.ViewResidentPrompt))
		color.Cyan(string(constants.ViewOfficersPrompt))
		color.Cyan(string(constants.LogoutPrompt))
		color.Blue(string(constants.ChoicePrompt))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handler.HandleViewResidents(*user)
		case "2":
			handler.HandleViewOfficers(*user)
		case "3":
			color.Red("Logging out...")
			return
		default:
			color.Red("Invalid choice, try again.")
		}
	}
}
