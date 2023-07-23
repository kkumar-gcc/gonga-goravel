package listeners

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
)

type SendResetPasswordNotification struct {
}

func (receiver *SendResetPasswordNotification) Signature() string {
	return "send_reset_password_notification"
}

func (receiver *SendResetPasswordNotification) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *SendResetPasswordNotification) Handle(args ...any) error {
	token := args[0].(string)
	email := args[1].(string)
	err := facades.Mail().To([]string{email}).
		Content(mail.Content{Subject: "Reset Password Notification", Html: receiver.mailTemplate(token, email)}).
		Send()
	if err != nil {
		return err
	}
	return nil
}

func (receiver *SendResetPasswordNotification) resetUrl(token string, email string) string {
	url := facades.Config().Env("APP_URL", "http://localhost")
	port := facades.Config().Env("APP_PORT", "3000")
	if url != "" && port != "" {
		return url.(string) + ":" + port.(string) + "/reset-password?token=" + token + "&email=" + email
	}
	return ""
}

func (receiver *SendResetPasswordNotification) mailTemplate(token string, email string) string {
	return `<html>
	<head>
		<title>Reset Password Notification</title>
	</head>
	<body>
		<p>Hi,</p>
		<p>We received a request to reset the password for your account.</p>
		<p>Click the button below to reset it.</p>
		<p><a href="` + receiver.resetUrl(token, email) + `">Reset Password</a></p>
		<p>If you did not request a password reset, please ignore this email or reply to let us know.</p>
		<p>Thanks,</p>
	</body>
</html>`
}
