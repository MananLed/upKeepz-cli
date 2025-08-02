package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/fatih/color"
	"github.com/common-nighthawk/go-figure"
)

func main(){

	userRepo := &repository.UserRepository{}
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	reader := bufio.NewReader(os.Stdin)

	for{
		myFigure := figure.NewColorFigure("UpKeepz","", "green", false)
		fmt.Println("🛠️                                                 🛠️")
		myFigure.Print()
		fmt.Println("🛠️                                                 🛠️")
		
		color.Cyan("1. Sign Up\n")
		color.Cyan("2. Login\n")
		color.Cyan("3. Exit\n\n")
		color.Blue("Enter your choice:-")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice{
		case "1":
			userHandler.SignUp()
		case "2":
			userHandler.Login()
		case "3":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}