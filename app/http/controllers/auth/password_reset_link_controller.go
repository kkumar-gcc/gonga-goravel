package auth

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/helpers"
)

type PasswordResetLinkController struct {
	//Dependent services
}

func NewPasswordResetLinkController() *PasswordResetLinkController {
	return &PasswordResetLinkController{
		//Inject services
	}
}

func (r *PasswordResetLinkController) Index(ctx http.Context) {
}

func (r *PasswordResetLinkController) Store(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"email": "required|email",
	})
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if validator.Fails() {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, validator.Errors().All())
		return
	}
	email := ctx.Request().Input("email")
	if err := helpers.SendResetLinkEmail(email); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "unable to send reset link")
		return
	}

	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "reset link sent",
	})
}
