package Auth

import (
	"github.com/goravel/framework/contracts/http"
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
	}

	//sendResetLinkEmail(ctx.Request().Query("email"))
	//
	//ctx.Response().Status(http.StatusNoContent)
}
