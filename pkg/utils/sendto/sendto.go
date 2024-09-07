package sendto

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/phongnd2802/go-backend-api/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MINE-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOTP(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{Address: from, Name: "test"},
		To: to,
		Subject: "OTP Verification",
		Body: fmt.Sprintf("Your OTP is %s. Please enter it to active your account.", otp),
	}

	messageEmail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", global.Config.SMTP.Username, global.Config.SMTP.Password, global.Config.SMTP.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", global.Config.SMTP.Host, global.Config.SMTP.Port), auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}

	return nil
}
