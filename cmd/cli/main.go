package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/MananLed/upKeepz-cli/internal/repository"
	"github.com/MananLed/upKeepz-cli/internal/service"
)

func main(){

	userRepo := &repository.UserRepository{}
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Println("\n======= UpKeepz =======")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("Enter your choice: ")

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