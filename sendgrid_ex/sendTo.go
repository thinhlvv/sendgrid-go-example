package sendgrid_ex

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Sendgrid struct {
	APIKey            string
	SenderName        string
	SenderEmail       string
	Subject           string
	Content           string
	Template          string
	TemplateName      string
	CreateTemplateUri string
}

const (
	SendgridHost = "https://api.sendgrid.com"
	SendUri      = "/v3/mail/send"
)

func NewEmail() *Sendgrid {
	return new(Sendgrid)
}

func (sg *Sendgrid) SendToPerson(receiver *mail.Email) {
	name := sg.SenderName
	address := sg.SenderEmail
	sender := mail.NewEmail(name, address)

	content := mail.NewContent("text/plain", sg.Content)
	m := mail.NewV3MailInit(sender, sg.Subject, receiver, content)

	request := sendgrid.GetRequest(sg.APIKey, SendUri, SendgridHost)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func (sg *Sendgrid) SendToPeople(emails ...*mail.Email) {
	name := sg.SenderName
	address := sg.SenderEmail
	sender := mail.NewEmail(name, address)

	content := mail.NewContent("text/plain", sg.Content)
	m := mail.NewV3MailInit(sender, sg.Subject, emails[0], content)

	m.Personalizations[0].AddTos(emails[1:]...)

	request := sendgrid.GetRequest(sg.APIKey, SendUri, SendgridHost)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
