package listeners

import (
	"github.com/goravel/framework/contracts/event"
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

	return nil
}
