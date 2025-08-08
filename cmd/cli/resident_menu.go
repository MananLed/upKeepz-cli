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

func ShowResidentDashboard(ctx context.Context, user *model.User, uHandler *handlers.UserHandler, sHandler *handlers.ServiceRequestHandler,
	nHandler *handlers.NoticeHandler, fHandler *handlers.FeedbackHandler, iHandler *handlers.InvoiceHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(constants.ResidentEmogiPrompt)
		myFigure := figure.NewColorFigure("Resident", "", "green", false)
		myFigure.Print()
		fmt.Println(constants.ResidentEmogiPrompt)

		color.Cyan("1. Manage Service Requests")
		color.Cyan("2. View Notices")
		color.Cyan("3. Manage Feedback")
		color.Cyan("4. View Invoices")
		color.Cyan("5. Manage Profile")
		color.Red("6. Logout")
		fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			for {
				color.Cyan("1." + string(constants.CreateServiceRequestPrompt))
				color.Cyan("2." + string(constants.RescheduleServiceRequestPrompt))
				color.Cyan("3." + string(constants.CancelServiceRequestPrompt))
				color.Cyan("4." + string(constants.GetApprovedServiceRequestPrompt))
				color.Cyan("5." + string(constants.GetPendingServiceRequestPrompt))
				color.Cyan("6. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					sHandler.BookServiceRequest(ctx)
				case "2":
					sHandler.RescheduleServiceRequest(ctx)
				case "3":
					sHandler.CancelServiceRequest(ctx)
				case "4":
					sHandler.GetApprovedServiceRequests(ctx)
				case "5":
					sHandler.GetPendingServiceRequests(ctx)
				case "6":
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
				color.Cyan("1." + string(constants.GetNoticePrompt))
				color.Cyan("2." + string(constants.GetNoticesByMonthYear))
				color.Cyan("3." + string(constants.GetNoticesByYear))
				color.Cyan("4. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					nHandler.GetNotices()
				case "2":
					nHandler.GetNoticesByMonthYear()
				case "3":
					nHandler.GetNoticesByYear()
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

		case "3":
			for {
				color.Cyan("1." + string(constants.GetFeedbackForResident))
				color.Cyan("2." + string(constants.GiveFeedbackPrompt))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					fHandler.GetFeebacksOfResident(ctx)
				case "2":
					fHandler.IssueFeedback(ctx)
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

		case "4":
			for {
				color.Cyan("1." + string(constants.SearchAInvoice))
				color.Cyan("2." + string(constants.ListInvoicesOfAYear))
				color.Cyan("3. Exit")
				fmt.Print(color.BlueString(string(constants.ChoicePrompt)))

				ch, _ := reader.ReadString('\n')
				ch = strings.TrimSpace(ch)
				exit := false

				switch ch {
				case "1":
					iHandler.GetInvoiceByMonthAndYear()
				case "2":
					iHandler.GetInvoicesByYear()
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

		case "5": 
			for {
				color.Cyan("1." + string(constants.UpdateProfilePrompt))
				color.Cyan("2." + string(constants.ChangePasswordPrompt))
				color.Cyan("3." + string(constants.ViewProfilePrompt))
				color.Red("4." + string(constants.DeleteProfilePrompt))
				color.Cyan("5. Exit")
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
					uHandler.ViewProfile(user)
				case "4":
					uHandler.DeleteProfile(ctx)
					return
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

		case "6": 
			color.Red("Logging out...")
			return

		default:
			color.Red("Invalid choice, try again.")
		}
	}
}
