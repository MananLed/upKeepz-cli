package constants 

type Prompt string
type Path string

const(
	ViewResidentPrompt Prompt = "1. View All Residents"
	ViewOfficersPrompt Prompt = "2. View All Officers"
	LogoutPrompt Prompt  = "3. Logout"
	SignUpPrompt Prompt = "1. Sign Up"
	LoginPrompt Prompt = "2. Login"
	ExitPrompt Prompt = "3. Exit\n\n"
	FirstNamePrompt Prompt = "First Name"
	MiddleNamePrompt Prompt = "Middle Name (optional):"
	LastNamePrompt Prompt = "Last Name"
	EmailPrompt Prompt = "Email"
	MobilePrompt Prompt = "Mobile Number"
	IDPrompt Prompt = "ID (Username)"
	PasswordPrompt Prompt = "Password: "
	ConfirmPasswordPrompt Prompt = "Confirm Password: "
	RolePrompt Prompt= "Role (Admin / MaintenanceOfficer / FlatResident)"
	ChoicePrompt Prompt = "Enter your choice:-"
	SignUpEmogiPrompt Prompt = "📝                                                📝"
	LoginEmogiPrompt Prompt = "🧑‍💻                                          🧑‍💻"
	AppEmogiPrompt Prompt = "🛠️                                                 🛠️"
	AdminEmogiPrompt Prompt = "🔐                                       🔐"
	
	MainDataPath Path = "../../data"
	UserDataPath Path = MainDataPath + "/users.json" 
)