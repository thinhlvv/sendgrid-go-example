package main

import (
	"os"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gitlab.com/sendgrid-go-example/sendgrid_ex"
)

func main() {
	mailSg := sendgrid_ex.NewEmail()

	mailSg.APIKey = os.Getenv("SENDGRID_API_KEY")
	mailSg.SenderEmail = "senderEmail@gmail.com"
	mailSg.Content = "This is a examle email from github.com/thinhlvv/sendgrid-go-example"
	mailSg.Subject = "Example"
	receiver1 := mail.Email{Name: "Receiver1", Address: "receiver1@gmail.com"}
	receiver2 := mail.Email{Name: "Receiver2", Address: "receiver2@gmail.com"}

	// mailSg.SendToPerson(&receiver1)
	mailSg.SendToPeople(&receiver1, &receiver2)
}
