package handlers

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

func PromptRequired(label string, reader *bufio.Reader) string {
	for {
		color.Yellow("%s: ", label)
		input, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(input)

		if trimmed == "" {
			color.Red("%s is compulsory", label)
			continue
		}
		return trimmed
	}
}

func ValidateMobileNumber(mobile string) bool{
	mobile = strings.TrimSpace(mobile)

	if len(mobile) != 10 {
		color.Red("Mobile number must be exactly 10 digits.")
		return false
	}

	if mobile[0] == '0' {
		color.Red("Mobile number cannot start with 0.")
		return false
	}

	if !isAllDigits(mobile) {
		color.Red("Mobile number must contain only digits.")
		return false
	}

	return true
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func (h *UserHandler) SignUp(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("📝                                                📝")
	myFigure := figure.NewColorFigure("Sign Up","", "purple", false)
	myFigure.Print()
	fmt.Println("📝                                                📝")

	firstName := PromptRequired("First Name", reader)

	color.Yellow("Middle Name (optional): ")
	middleName, _ := reader.ReadString('\n')

	lastName := PromptRequired("Last Name", reader)

	color.Yellow("Email: ")
	email, _ := reader.ReadString('\n')

	mobile := PromptRequired("Mobile Number", reader)
	for{
		if ValidateMobileNumber(mobile) {
			break
		}
		mobile = PromptRequired("Mobile Number", reader)
	}

	id := PromptRequired("ID (Username)", reader)

	var passwordStr string
	for {
		color.Yellow("Password")
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		passwordStr = string(password)

		if passwordStr == "" {
			color.Red("Password is compulsory")
			continue
		}
		break
	}

	roleStr := PromptRequired("Role (Admin / MaintenanceOfficer / FlatResident)", reader)
	for{
		if roleStr == "Admin" || roleStr == "MaintenanceOfficer" || roleStr == "FlatResident"{
			break
		}
		color.Red("Invalid Role")
		roleStr = PromptRequired("Role (Admin / MaintenanceOfficer / FlatResident)", reader)
	}

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

	fmt.Println("🧑‍💻                                          🧑‍💻")
	myFigure := figure.NewColorFigure("Login","", "blue", false)
	myFigure.Print()
	fmt.Println("🧑‍💻                                          🧑‍💻")

	var id string
	for {
		color.Yellow("ID (username): ")
		input, _ := reader.ReadString('\n')
		id = strings.TrimSpace(input)

		if id == "" {
			color.Red("ID is required")
			continue
		}
		break
	}

	var passwordStr string
	for {
		color.Yellow("Password")
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		passwordStr = string(password)

		if passwordStr == "" {
			color.Red("Password is required")
			continue
		}
		break
	}

	id = strings.TrimSpace(id)
	passwordStr = strings.TrimSpace(passwordStr)

	user, err := h.UserService.Login(id, passwordStr)

	if err != nil{
		fmt.Println("Login failed:", err)
		return nil
	}

	fmt.Println("Login successful!! Welcome,", user.FirstName)
	return user
}