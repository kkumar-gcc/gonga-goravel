package events

import "github.com/goravel/framework/contracts/event"

type EmailVerificationNotification struct {
}

func (receiver *EmailVerificationNotification) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
