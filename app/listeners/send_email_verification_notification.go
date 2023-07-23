package listeners

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
	"strconv"
)

type SendEmailVerificationNotification struct {
}

func (receiver *SendEmailVerificationNotification) Signature() string {
	return "send_email_verification_notification"
}

func (receiver *SendEmailVerificationNotification) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *SendEmailVerificationNotification) Handle(args ...any) error {
	id := args[0].(uint)
	email := args[1].(string)
	err := facades.Mail().To([]string{email}).
		Content(mail.Content{Subject: "Verify Email Address", Html: receiver.mailTemplate(id, email)}).
		Send()
	if err != nil {
		return err
	}
	return nil
}

func (receiver *SendEmailVerificationNotification) verificationUrl(id uint, email string) string {
	url := facades.Config().GetString("app.frontend_url", "http://localhost:3000")
	hash, err := facades.Crypt().EncryptString(email)
	if err != nil {
		return ""
	}
	return url + "/verify-email/" + strconv.Itoa(int(id)) + "/" + hash
}

func (receiver *SendEmailVerificationNotification) mailTemplate(id uint, email string) string {
	return `<html>
	<head>
		<title>Verify Email Address</title>
	</head>
	<body>
		<p>Hi,</p>
		<p>Please click the button below to verify your email address</p>
		<p><a href="` + receiver.verificationUrl(id, email) + `">Verify Email Address</a></p>
		<p>If you did not create an account, no further action is required.</p>
		<p>Thanks,</p>
	</body>
</html>`
}
