package Auth

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

		return
	}
	if validator.Fails() {
		ctx.Response().Json(http.StatusUnprocessableEntity, validator.Errors().All())
		return
	}
	email := ctx.Request().Input("email")
	if err := helpers.SendResetLinkEmail(email); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Unable to send reset link",
		})
		return
	}

	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Reset link sent",
	})
}
