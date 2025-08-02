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
)

func ShowAdminDashboard(user *model.User, handler *handlers.SocietyHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("🔐                                       🔐")
		myFigure := figure.NewColorFigure("Admin","", "green", false)
		myFigure.Print()
		fmt.Println("🔐                                       🔐")
		color.Cyan("1. View All Residents")
		color.Cyan("2. View All Officers")
		color.Cyan("3. Logout")
		color.Blue("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handler.HandleViewResidents(*user)
		case "2":
			handler.HandleViewOfficers(*user)
		case "3":
			fmt.Println("Logging out...\n")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}