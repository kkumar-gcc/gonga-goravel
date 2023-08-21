package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
)

type EmailVerificationNotificationController struct {
	//Dependent services
}

func NewEmailVerificationNotificationController() *EmailVerificationNotificationController {
	return &EmailVerificationNotificationController{
		//Inject services
	}
}

func (r *EmailVerificationNotificationController) Index(ctx http.Context) {
}

func (r *EmailVerificationNotificationController) Store(ctx http.Context) {
	var user models.User
	err := facades.Auth().User(ctx, &user) // Must point
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if user.HasVerifiedEmail() {
		url := facades.Config().GetString("app.frontend_url")
		ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")
		return
	}

	if err := helpers.SendEmailVerificationLink(ctx); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Json(http.StatusNoContent, http.Json{
		"message": "verification link sent",
	})
}
