package helpers

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"goravel/app/events"
	"goravel/app/models"
)

func SendEmailVerificationLink(ctx http.Context) error {
	var user models.User
	err := facades.Auth().User(ctx, &user) // Must point
	if err != nil {
		return err
	}
	err = facades.Event().Job(&events.EmailVerificationNotification{}, []event.Arg{
		{Type: "uint", Value: user.ID},
		{Type: "string", Value: user.Email},
	}).Dispatch()
	if err != nil {
		return err
	}
	return nil
}

func MarkEmailAsVerified(user models.User) error {
	user.EmailVerifiedAt = carbon.DateTime{
		Carbon: carbon.Now(),
	}
	if _, err := facades.Orm().Query().Model(&models.User{}).Update(&user); err != nil {
		return err
	}
	return nil
}
