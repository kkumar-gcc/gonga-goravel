package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type EmailVerificationRequest struct {
}

func (r *EmailVerificationRequest) Authorize(ctx http.Context) error {
	var user models.User
	err := facades.Auth().User(ctx, &user) // Must point
	if err != nil {
		return err
	}
	if user.HasVerifiedEmail() {
		url := facades.Config().GetString("app.frontend_url")
		ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")
		return nil
	}

	return nil
}

func (r *EmailVerificationRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *EmailVerificationRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *EmailVerificationRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *EmailVerificationRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
