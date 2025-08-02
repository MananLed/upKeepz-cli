package handlers

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/fatih/color"
	"github.com/common-nighthawk/go-figure"
	"gitlab.com/david_mbuvi/go_asterisks"
)

type UserHandler struct{
	UserService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler{
	return &UserHandler{
		UserService: us,
	}
}

func (h *UserHandler) SignUp(){
	reader := bufio.NewReader(os.Stdin)

	myFigure := figure.NewColorFigure("Sign Up","", "purple", false)
	myFigure.Print()

	color.Yellow("First Name: ")
	firstName, _ := reader.ReadString('\n')

	color.Yellow("Middle Name (optional): ")
	middleName, _ := reader.ReadString('\n')

	color.Yellow("Last Name: ")
	lastName, _ := reader.ReadString('\n')

	color.Yellow("Email: ")
	email, _ := reader.ReadString('\n')

	color.Yellow("Mobile Number: ")
	mobile, _ := reader.ReadString('\n')

	color.Yellow("ID (username): ")
	id, _ := reader.ReadString('\n')

	color.Yellow("Password: ")
	password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	passwordStr := string(password)

	color.Yellow("Role (Admin / MaintenanceOfficer / FlatResident): ")
	roleStr, _ := reader.ReadString('\n')

	user := model.User{
		FirstName:    strings.TrimSpace(firstName),
		MiddleName:   strings.TrimSpace(middleName),
		LastName:     strings.TrimSpace(lastName),
		Email:        strings.TrimSpace(email),
		MobileNumber: strings.TrimSpace(mobile),
		ID:           strings.TrimSpace(id),
		Password:     strings.TrimSpace(passwordStr),
		Role:         model.ParseRole(strings.TrimSpace(roleStr)),
	}

	err := h.UserService.SignUp(user)

	if err != nil{
		fmt.Println("Sign up failed: ", err)
	}else {
		fmt.Println("User signed up successfully!!")
	}
}

func (h *UserHandler) Login() *model.User {
	reader := bufio.NewReader(os.Stdin)
	myFigure := figure.NewColorFigure("Login","", "blue", false)
	myFigure.Print()

	color.Yellow("ID (username): ")
	id, _ := reader.ReadString('\n')

	color.Yellow("Password: ")
	password, _ := reader.ReadString('\n')

	id = strings.TrimSpace(id)
	password = strings.TrimSpace(password)

	user, err := h.UserService.Login(id, password)

	if err != nil{
		fmt.Println("Login failed:", err)
		return nil
	}

	fmt.Println("Login successful!! Welcome,", user.FirstName)
	return user
}