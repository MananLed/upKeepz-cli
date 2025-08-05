// package handlers

// import(
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"unicode"
// 	"github.com/MananLed/upKeepz-cli/internal/model"
// 	"github.com/MananLed/upKeepz-cli/internal/service"
// 	"github.com/MananLed/upKeepz-cli/constants"
// 	"github.com/fatih/color"
// 	"github.com/common-nighthawk/go-figure"
// 	"gitlab.com/david_mbuvi/go_asterisks"
// )

// type UserHandler struct{
// 	UserService *service.UserService
// }

// func NewUserHandler(us *service.UserService) *UserHandler{
// 	return &UserHandler{
// 		UserService: us,
// 	}
// }

// func PromptRequired(label string, reader *bufio.Reader) string {
// 	for {
// 		color.Yellow("%s: ", label)
// 		input, _ := reader.ReadString('\n')
// 		trimmed := strings.TrimSpace(input)

// 		if trimmed == "" {
// 			color.Red("%s is compulsory", label)
// 			continue
// 		}
// 		return trimmed
// 	}
// }

// func ValidateMobileNumber(mobile string) bool{
// 	mobile = strings.TrimSpace(mobile)

// 	if len(mobile) != 10 {
// 		color.Red("Mobile number must be exactly 10 digits.")
// 		return false
// 	}

// 	if mobile[0] == '0' {
// 		color.Red("Mobile number cannot start with 0.")
// 		return false
// 	}

// 	if !isAllDigits(mobile) {
// 		color.Red("Mobile number must contain only digits.")
// 		return false
// 	}

// 	return true
// }

// func isAllDigits(s string) bool {
// 	for _, r := range s {
// 		if !unicode.IsDigit(r) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (h *UserHandler) SignUp(){
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println(constants.SignUpEmogiPrompt)
// 	myFigure := figure.NewColorFigure("Sign Up","", "purple", false)
// 	myFigure.Print()
// 	fmt.Println(constants.SignUpEmogiPrompt)

// 	firstName := PromptRequired(string(constants.FirstNamePrompt), reader)

// 	color.Yellow(string(constants.MiddleNamePrompt))
// 	middleName, _ := reader.ReadString('\n')

// 	lastName := PromptRequired(string(constants.LastNamePrompt), reader)

// 	email := PromptRequired(string(constants.EmailPrompt), reader)

// 	mobile := PromptRequired(string(constants.MobilePrompt), reader)
// 	for{
// 		if ValidateMobileNumber(mobile) {
// 			break
// 		}
// 		mobile = PromptRequired(string(constants.MobilePrompt), reader)
// 	}

// 	id := PromptRequired(string(constants.IDPrompt), reader)

// 	var passwordStr string
// 	for {
// 		color.Yellow(string(constants.PasswordPrompt))
// 		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
// 		passwordStr = string(password)

// 		if passwordStr == "" {
// 			color.Red("Password is compulsory")
// 			continue
// 		}
// 		break
// 	}

// 	var confirmPasswordStr string
// 	for {
// 		color.Yellow(string(constants.ConfirmPasswordPrompt))
// 		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
// 		confirmPasswordStr = string(password)

// 		if passwordStr != confirmPasswordStr {
// 			color.Red("Password does'nt match")
// 			continue
// 		}
// 		break
// 	}

// 	roleStr := PromptRequired(string(constants.RolePrompt), reader)
// 	for{
// 		if roleStr == "Admin" || roleStr == "MaintenanceOfficer" || roleStr == "FlatResident"{
// 			break
// 		}
// 		color.Red("Invalid Role")
// 		roleStr = PromptRequired(string(constants.RolePrompt), reader)
// 	}

// 	user := model.User{
// 		FirstName:    strings.TrimSpace(firstName),
// 		MiddleName:   strings.TrimSpace(middleName),
// 		LastName:     strings.TrimSpace(lastName),
// 		Email:        strings.TrimSpace(email),
// 		MobileNumber: strings.TrimSpace(mobile),
// 		ID:           strings.TrimSpace(id),
// 		Password:     strings.TrimSpace(passwordStr),
// 		Role:         model.ParseRole(strings.TrimSpace(roleStr)),
// 	}

// 	err := h.UserService.SignUp(user)

// 	if err != nil{
// 		color.Red("Sign up failed: %v", err)
// 	}else {
// 		color.Green("User signed up successfully!!")
// 	}
// }

// func (h *UserHandler) Login() *model.User {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println(constants.LoginEmogiPrompt)
// 	myFigure := figure.NewColorFigure("Login","", "blue", false)
// 	myFigure.Print()
// 	fmt.Println(constants.LoginEmogiPrompt)

// 	var id string
// 	for {
// 		color.Yellow("ID (username): ")
// 		input, _ := reader.ReadString('\n')
// 		id = strings.TrimSpace(input)

// 		if id == "" {
// 			color.Red("ID is required")
// 			continue
// 		}
// 		break
// 	}

// 	var passwordStr string
// 	for {
// 		color.Yellow("Password")
// 		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
// 		passwordStr = string(password)

// 		if passwordStr == "" {
// 			color.Red("Password is required")
// 			continue
// 		}
// 		break
// 	}

// 	id = strings.TrimSpace(id)
// 	passwordStr = strings.TrimSpace(passwordStr)

// 	user, err := h.UserService.Login(id, passwordStr)

// 	if err != nil{
// 		color.Red("Login failed: %v", err)
// 		return nil
// 	}

// 	color.Green("Login successful!! Welcome,", user.FirstName)
// 	return user
// }

package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"gitlab.com/david_mbuvi/go_asterisks"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
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

func ValidateMobileNumber(mobile string) bool {
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

func (h *UserHandler) SignUp() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(constants.SignUpEmogiPrompt)
	myFigure := figure.NewColorFigure("Sign Up", "", "purple", false)
	myFigure.Print()
	fmt.Println(constants.SignUpEmogiPrompt)

	firstName := PromptRequired(string(constants.FirstNamePrompt), reader)

	color.Yellow(string(constants.MiddleNamePrompt))
	middleName, _ := reader.ReadString('\n')


	lastName := PromptRequired(string(constants.LastNamePrompt), reader)

	email := PromptRequired(string(constants.EmailPrompt), reader)

	mobile := PromptRequired(string(constants.MobilePrompt), reader)
	for {
		if ValidateMobileNumber(mobile) {
			break
		}
		mobile = PromptRequired(string(constants.MobilePrompt), reader)
	}

	id := PromptRequired(string(constants.IDPrompt), reader)

	var passwordStr string
	for {
		color.Yellow(string(constants.PasswordPrompt))
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		passwordStr = string(password)

		if passwordStr == "" {
			color.Red("Password is compulsory")
			continue
		}
		break
	}

	var confirmPasswordStr string
	for {
		color.Yellow(string(constants.ConfirmPasswordPrompt))
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		confirmPasswordStr = string(password)

		if passwordStr != confirmPasswordStr {
			color.Red("Password doesn't match")
			continue
		}
		break
	}

	roleStr := PromptRequired(string(constants.RolePrompt), reader)
	for {
		if roleStr == "Admin" || roleStr == "MaintenanceOfficer" || roleStr == "FlatResident" {
			break
		}
		color.Red("Invalid Role")
		roleStr = PromptRequired(string(constants.RolePrompt), reader)
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

	if err != nil {
		color.Red("Sign up failed: %v", err)
	} else {
		color.Green("User signed up successfully!!")
	}
}

func (h *UserHandler) Login() *model.User {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(constants.LoginEmogiPrompt)
	myFigure := figure.NewColorFigure("Login", "", "blue", false)
	myFigure.Print()
	fmt.Println(constants.LoginEmogiPrompt)

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


	user, err := h.UserService.Login(id, passwordStr)
	if err != nil {
		color.Red("Login failed: %v", err)
		return nil
	}

	color.Green("Login successful!! Welcome, %s", user.FirstName)

	// Store user in context
	return user 
}
