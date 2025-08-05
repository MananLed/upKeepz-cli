package constants 

type Prompt string
type Path string

const(
	ViewResidentPrompt Prompt = "View All Residents"
	ViewOfficersPrompt Prompt = "View All Officers"
	LogoutPrompt Prompt  = "Logout"
	SignUpPrompt Prompt = "Sign Up"
	LoginPrompt Prompt = "Login"
	ExitPrompt Prompt = "Exit\n\n"
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
	DeleteResidentPrompt Prompt = "Delete Resident Credentials"
	DeleteOfficerPrompt Prompt = "Delete Officer Credentials"
	IssueNoticePrompt Prompt = "Issue Notice"
	GetNoticePrompt Prompt = "Get Notices"
	GetNoticeByID Prompt = "Get Notice By ID"
	MainDataPath Path = "../../data"
	UserDataPath Path = MainDataPath + "/users.json" 
	NoticeDataPath Path = MainDataPath + "/notices.json"
)