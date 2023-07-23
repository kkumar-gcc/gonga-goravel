package helpers

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"goravel/app/events"
	"goravel/app/models"
)

func SendResetLinkEmail(email string) error {
	var user models.User
	if err := facades.Orm().Query().Where("email", email).First(&user); err != nil {
		return err
	}
	token, err := CreateToken(user)
	if err != nil {
		return err
	}
	err = facades.Event().Job(&events.ResetPasswordNotification{}, []event.Arg{
		{Type: "string", Value: token},
		{Type: "string", Value: user.Email},
	}).Dispatch()
	if err != nil {
		return err
	}
	return nil
}

func CreateToken(user models.User) (string, error) {
	email := user.Email

	if err := DeleteExisting(user); err != nil {
		return "", err
	}

	token, err := facades.Hash().Make(email)
	if err != nil {
		return "", err
	}

	err = facades.Orm().Query().Model(&models.PasswordReset{}).Create(&models.PasswordReset{
		Email: email,
		Token: token,
		ExpiresAt: carbon.DateTime{
			Carbon: carbon.Now().AddMinutes(60),
		},
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func DeleteExisting(user models.User) error {
	var passwordReset models.PasswordReset
	if _, err := facades.Orm().Query().Model(&models.PasswordReset{}).Where("email", user.Email).Delete(&passwordReset); err != nil {
		return err
	}
	return nil
}

func PasswordReset(ctx http.Context) error {
	token := ctx.Request().Input("token")
	email := ctx.Request().Input("email")
	var passwordReset models.PasswordReset
	err := facades.Orm().Query().Model(&models.PasswordReset{}).Where("email", email).Where("token", token).First(&passwordReset)
	if err != nil {
		return err
	}
	password := ctx.Request().Input("password")
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		return err
	}
	_, err = facades.Orm().Query().Where("email", email).Update(map[string]interface{}{
		"password": hashedPassword,
	})
	if err != nil {
		return err
	}
	_, err = facades.Orm().Query().Model(&models.PasswordReset{}).Where("email", email).Delete(&passwordReset)
	if err != nil {
		return err
	}
	return nil
}
