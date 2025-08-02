package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/handlers"
)

func ShowAdminDashboard(user *model.User, handler *handlers.SocietyHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n🔐 Admin Dashboard")
		fmt.Println("1. View All Residents")
		fmt.Println("2. View All Officers")
		fmt.Println("3. Logout")
		fmt.Print("Choose an option: ")

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