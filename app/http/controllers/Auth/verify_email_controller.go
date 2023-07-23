package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	Auth2 "goravel/app/http/requests/Auth"
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
	var emailVerificationRequest Auth2.EmailVerificationRequest
	errors, err := ctx.Request().ValidateRequest(&emailVerificationRequest)
	if err != nil || errors != nil {
		ctx.Response().Json(http.StatusUnprocessableEntity, errors.All())
		return
	}

	if emailVerificationRequest.Authorize(ctx) != nil {
		return
	}
	var user models.User
	err = facades.Auth().User(ctx, &user) // Must point
	if err != nil {
		return
	}

	if helpers.MarkEmailAsVerified(user) != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	url := facades.Config().GetString("app.frontend_url")
	ctx.Response().Redirect(http.StatusFound, url+"/"+"?verified=1")

}
