package usecase

import (
	"fmt"
	"mini-socmed/internal/constant"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type (
	EmailSenderUsecase interface {
		SendEmail(subject, content, to string) error
	}
	emailSenderUsecase struct {
		name         string
		fromEmail    string
		fromPassword string
	}
)

func (e *emailSenderUsecase) SendEmail(subject string, content string, to string) error {
	mail := email.NewEmail()
	mail.From = fmt.Sprintf("%s <%s>", e.name, e.fromEmail)
	mail.Subject = subject
	mail.HTML = []byte(content)
	mail.To = []string{to}

	smtpAuth := smtp.PlainAuth("", e.fromEmail, e.fromPassword, constant.SmtpAuthAddress)
	if err := mail.Send(constant.SmtpServerAddress, smtpAuth); err != nil {
		return err
	}
	return nil
}

func NewEmailSenderUsecase(name string, fromEmail string, fromPassword string) EmailSenderUsecase {
	return &emailSenderUsecase{
		name:         name,
		fromEmail:    fromEmail,
		fromPassword: fromPassword,
	}

}
