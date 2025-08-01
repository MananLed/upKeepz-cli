package handlers

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
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

	fmt.Println("========= Sign Up =======")

	fmt.Print("First Name: ")
	firstName, _ := reader.ReadString('\n')

	fmt.Print("Middle Name (optional): ")
	middleName, _ := reader.ReadString('\n')

	fmt.Print("Last Name: ")
	lastName, _ := reader.ReadString('\n')

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Mobile Number: ")
	mobile, _ := reader.ReadString('\n')

	fmt.Print("ID (username): ")
	id, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')

	fmt.Print("Role (Admin / MaintenanceOfficer / FlatResident): ")
	roleStr, _ := reader.ReadString('\n')

	user := model.User{
		FirstName:    strings.TrimSpace(firstName),
		MiddleName:   strings.TrimSpace(middleName),
		LastName:     strings.TrimSpace(lastName),
		Email:        strings.TrimSpace(email),
		MobileNumber: strings.TrimSpace(mobile),
		ID:           strings.TrimSpace(id),
		Password:     strings.TrimSpace(password),
		Role:         model.ParseRole(strings.TrimSpace(roleStr)),
	}

	err := h.UserService.SignUp(user)

	if err != nil{
		fmt.Println("Sign up failed: ", err)
	}else {
		fmt.Println("User signed up successfully!!")
	}
}

func (h *UserHandler) Login(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======= Login =======")

	fmt.Print("ID (username): ")
	id, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')

	id = strings.TrimSpace(id)
	password = strings.TrimSpace(password)

	user, err := h.UserService.Login(id, password)

	if err != nil{
		fmt.Println("Login failed:", err)
		return
	}

	fmt.Println("Login successful!! Welcome,", user.FirstName)
}