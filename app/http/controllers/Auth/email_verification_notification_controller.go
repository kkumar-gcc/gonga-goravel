package Auth

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
		return
	}
	if user.HasVerifiedEmail() {
		url := facades.Config().GetString("app.frontend_url")
		ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")
		return
	}

	if err := helpers.SendEmailVerificationLink(ctx); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	ctx.Response().Json(http.StatusNoContent, http.Json{
		"message": "verification link sent",
	})
}
