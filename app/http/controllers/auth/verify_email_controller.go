package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
)

type VerifyEmailController struct {
	//Dependent services
}

func NewVerifyEmailController() *VerifyEmailController {
	return &VerifyEmailController{
		//Inject services
	}
}

func (r *VerifyEmailController) Index(ctx http.Context) {
	var user models.User
	err := facades.Auth().User(ctx, &user) // Must point
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	url := facades.Config().GetString("app.frontend_url")
	if user.HasVerifiedEmail() {
		ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")
		return
	}

	if err := helpers.MarkEmailAsVerified(user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")
}
