package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/fatih/color"
	"github.com/common-nighthawk/go-figure"
	"github.com/MananLed/upKeepz-cli/constants"
)

func main(){

	userRepo := &repository.UserRepository{}
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	societyRepo := &repository.SocietyRepository{}
    societyService := service.NewSocietyService(societyRepo)
    societyHandler := handlers.NewSocietyHandler(societyService)

	reader := bufio.NewReader(os.Stdin)

	for{
		myFigure := figure.NewColorFigure("UpKeepz","", "green", false)
		fmt.Println(constants.AppEmogiPrompt)
		myFigure.Print()
		fmt.Println(constants.AppEmogiPrompt)

		color.Cyan(string(constants.SignUpPrompt))
		color.Cyan(string(constants.LoginPrompt))
		color.Cyan(string(constants.ExitPrompt))
		color.Blue(string(constants.ChoicePrompt))

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice{
		case "1":
			userHandler.SignUp()
		case "2":
			user := userHandler.Login()
			if user == nil {continue}
			switch user.Role {
			case model.RoleAdmin:
				ShowAdminDashboard(user, societyHandler)
			default:
				color.Green("Logged in as", user.Role)
			}
		case "3":
			color.Red("Exit")
			return
		default:
			color.Red("Invalid choice. Please try again.")
		}
	}
}