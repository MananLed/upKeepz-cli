package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/service"
	"github.com/MananLed/upKeepz-cli/internal/utils"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"gitlab.com/david_mbuvi/go_asterisks"
	"golang.org/x/crypto/bcrypt"
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
		fmt.Print(color.YellowString(label) + ": ")
		input, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(input)
		trimmed = strings.TrimRight(trimmed, "\r\n")
		if trimmed == "" {
			color.Red("%s is compulsory", label)
			continue
		}
		return trimmed
	}
}

func ValidateMobileNumber(mobile string) bool {
	mobile = strings.TrimSpace(mobile)

	pattern := `^[6-9][0-9]{9}$`

	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		color.Red("Invalid email format.")
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	var hasLower, hasDigit, hasSpecial bool

	if len(password) < 12 {
		color.Red("Password must be at least 12 characters long.")
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasLower || !hasDigit || !hasSpecial {
		color.Red("Password must contain at least one lowercase letter, one digit, and one special character.")
		return false
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

	fmt.Print(color.YellowString(string(constants.MiddleNamePrompt)))
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

	id := email

	var passwordStr string
	for {
		fmt.Print(color.YellowString(string(constants.PasswordPrompt)))
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		passwordStr = string(password)

		if passwordStr == "" {
			color.Red("Password is compulsory")
			continue
		}

		if !ValidatePassword(passwordStr) {
			color.Red("Invalid password, enter again")
			continue
		}
		break
	}

	var confirmPasswordStr string
	for {
		fmt.Print(color.YellowString(string(constants.ConfirmPasswordPrompt)))
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		confirmPasswordStr = string(password)

		if passwordStr != confirmPasswordStr {
			color.Red("Password doesn't match")
			continue
		}
		break
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)

	roleStr := "FlatResident"

	user := model.User{
		FirstName:    strings.TrimSpace(firstName),
		MiddleName:   strings.TrimSpace(middleName),
		LastName:     strings.TrimSpace(lastName),
		Email:        strings.TrimSpace(email),
		MobileNumber: strings.TrimSpace(mobile),
		ID:           strings.TrimSpace(id),
		Password:     string(hashedPassword),
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
		fmt.Print(color.YellowString(string(constants.IDPrompt)))
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
		fmt.Print(color.YellowString(string(constants.PasswordPrompt)))
		password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		passwordStr = string(password)
		passwordStr = strings.TrimRight(passwordStr, "\r\n")
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
	return user
}

func (h *UserHandler) UpdateProfile(user *model.User) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.YellowString("Update First Name: "))
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)
	firstName = strings.TrimRight(firstName, "\r\n")
	if firstName != "" {
		user.FirstName = firstName
	}

	fmt.Print(color.YellowString("Update Middle Name: "))
	middleName, _ := reader.ReadString('\n')
	middleName = strings.TrimSpace(middleName)
	middleName = strings.TrimRight(middleName, "\r\n")
	if middleName != "" {
		user.MiddleName = middleName
	}

	fmt.Print(color.YellowString("Update Last Name: "))
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)
	lastName = strings.TrimRight(lastName, "\r\n")
	if lastName != "" {
		user.LastName = lastName
	}

	fmt.Print(color.YellowString("Update Mobile Number: "))
	mobile, _ := reader.ReadString('\n')
	mobile = strings.TrimSpace(mobile)
	mobile = strings.TrimRight(mobile, "\r\n")
	if mobile != "" && ValidateMobileNumber(mobile) {
		user.MobileNumber = mobile
	}

	fmt.Print(color.YellowString("Update Email/ID: "))
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	email = strings.TrimRight(email, "\r\n")
	if email != "" && ValidateEmail(email) {
		user.Email = email
		user.ID = email
	}

	// fmt.Print(color.YellowString("Enter Password: "))
	// password, _ := reader.ReadString('\n')
	// password = strings.TrimSpace(password)
	// password = strings.TrimRight(password, "\r\n")


	if err := h.UserService.UpdateProfile(*user); err != nil {
		color.Red("Failed to update profile: %v", err)
		return
	}

	color.Green("Profile updated successfully!")
}

func (h *UserHandler) ChangePassword(ctx context.Context) {

	fmt.Print(color.YellowString("Enter current password: "))
	oldPasswordBytes, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	oldPassword := string(oldPasswordBytes)
	oldPassword = strings.TrimRight(oldPassword, "\r\n")
	var newPassword, confirmNewPassword string

	for {
		fmt.Print(color.YellowString("Enter new password: "))
		passBytes, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		newPassword = string(passBytes)
		newPassword = strings.TrimRight(newPassword, "\r\n")

		if !ValidatePassword(newPassword) {
			color.Red("Password must have at least 12 characters, one lowercase, one digit, and one special character.")
			continue
		}

		fmt.Print(color.YellowString("Confirm new password: "))
		confirmBytes, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		confirmNewPassword = string(confirmBytes)
		confirmNewPassword = strings.TrimRight(confirmNewPassword, "\r\n")

		if newPassword != confirmNewPassword {
			color.Red("Passwords do not match.")
			continue
		}
		break
	}

	err := h.UserService.ChangePassword(ctx, oldPassword, newPassword)
	if err != nil {
		color.Red("Failed to change password: %v", err)
		return
	}

	color.Green("Password changed successfully!")
}

func (h *UserHandler) CreateOfficer(ctx context.Context) {
	currentUser, err := utils.GetUserFromContext(ctx)
	if err != nil || (currentUser.Role != model.RoleAdmin && currentUser.Role != model.RoleOfficer) {
		color.Red("Unauthorized: Only Admin or Officer can add a new officer.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	email := PromptRequired("Officer Email (used as ID)", reader)

	fmt.Print(color.YellowString("Set temporary password for the officer: "))
	password, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	passwordStr := string(password)
	passwordStr = strings.TrimRight(passwordStr, "\r\n")
	password = []byte(passwordStr)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	newOfficer := model.User{
		Email:    email,
		ID:       email,
		Password: string(hashedPassword),
		Role:     model.RoleOfficer,
	}

	if err := h.UserService.SignUp(newOfficer); err != nil {
		color.Red("Error creating officer: %v", err)
	} else {
		color.Green("Officer created successfully.")
	}
}
