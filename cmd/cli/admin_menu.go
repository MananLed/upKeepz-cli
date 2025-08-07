package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/handlers"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func ShowAdminDashboard(ctx context.Context, user *model.User, uHandler *handlers.UserHandler, socHandler *handlers.SocietyHandler, cHandler *handlers.CredentialHandler,
	nHandler *handlers.NoticeHandler, fHandler *handlers.FeedbackHandler, iHandler *handlers.InvoiceHandler, sHandler *handlers.ServiceRequestHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(constants.AdminEmogiPrompt)
		myFigure := figure.NewColorFigure("Admin", "", "green", false)
		myFigure.Print()
		fmt.Println(constants.AdminEmogiPrompt)

		color.Cyan("1. Manage Residents")
		color.Cyan("2. Manage Officers")
		color.Cyan("3. Manage Service Requests")
		color.Cyan("4. Manage Notices")
		color.Cyan("5. Manage Feedback")
		color.Cyan("6. Manage Invoices")
		color.Cyan("7. Manage Profile")
		color.Cyan("8. Add New Officer")
		color.Red("9. Logout")
		fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			for {
				color.Cyan("1." + string(constants.ViewResidentPrompt))
				color.Cyan("2." + string(constants.DeleteResidentPrompt))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					socHandler.HandleViewResidents(ctx)
				case "2":
					cHandler.DeleteResident(ctx)
				case "3":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "2":
			for {
				color.Cyan("1." + string(constants.ViewOfficersPrompt))
				color.Cyan("2." + string(constants.DeleteOfficerPrompt))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					socHandler.HandleViewOfficers(ctx)
				case "2":
					cHandler.DeleteOfficer(ctx)
				case "3":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "3":
			for {
				color.Cyan("1." + string(constants.GetPendingServiceRequestPrompt))
				color.Cyan("2." + string(constants.GetApprovedServiceRequestPrompt))
				color.Cyan("3." + string(constants.ApproveServiceRequestPrompt))
				color.Cyan("4. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					sHandler.ViewPendingRequestsByServiceType(ctx)
				case "2":
					sHandler.ViewApprovedRequestsByServiceType(ctx)
				case "3":
					sHandler.ApproveRequest(ctx)
				case "4":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "4":
			for {
				color.Cyan("1." + string(constants.IssueNoticePrompt))
				color.Cyan("2." + string(constants.GetNoticePrompt))
				color.Cyan("3." + string(constants.GetNoticesByMonthYear))
				color.Cyan("4." + string(constants.GetNoticesByYear))
				color.Cyan("5. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					nHandler.IssueNotice(ctx)
				case "2":
					nHandler.GetNotices()
				case "3":
					nHandler.GetNoticesByMonthYear()
				case "4":
					nHandler.GetNoticesByYear()
				case "5":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "5":
			for {
				color.Cyan("1." + string(constants.GetAllFeedbackPrompt))
				color.Cyan("2." + string(constants.GetFeedbacksByID))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					fHandler.GetFeedbacks(ctx)
				case "2":
					fHandler.GetFeebacksByResidentID(ctx)
				case "3":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "6":
			for {
				color.Cyan("1." + string(constants.IssueInvoice))
				color.Cyan("2." + string(constants.SearchAInvoice))
				color.Cyan("3." + string(constants.ListInvoicesOfAYear))
				color.Cyan("4. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					iHandler.IssueInvoice(ctx)
				case "2":
					iHandler.GetInvoiceByMonthAndYear()
				case "3":
					iHandler.GetInvoicesByYear()
				case "4":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "7":
			for {
				color.Cyan("1." + string(constants.UpdateProfilePrompt))
				color.Cyan("2." + string(constants.ChangePasswordPrompt))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))
				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					uHandler.UpdateProfile(user)
				case "2":
					uHandler.ChangePassword(ctx)
				case "3":
					color.Red("Exit")
					exit = true
				default:
					color.Red("Invalid choice, try again.")
				}
				if exit {
					break
				}
			}

		case "8":
			uHandler.CreateOfficer(ctx)

		case "9":
			color.Red("Logging out...")
			return

		default:
			color.Red("Invalid choice, try again.")
		}
	}
}

