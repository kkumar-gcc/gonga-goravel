package events

import "github.com/goravel/framework/contracts/event"

type ResetPasswordNotification struct {
}

func (receiver *ResetPasswordNotification) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
