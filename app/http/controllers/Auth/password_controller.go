package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/helpers"
)

type PasswordController struct {
	//Dependent services
}

func NewPasswordController() *PasswordController {
	return &PasswordController{
		//Inject services
	}
}

func (r *PasswordController) Index(ctx http.Context) {
}

func (r *PasswordController) Store(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"token":    "required",
		"email":    "required|email",
		"password": "required|min_len:8|eq_field:password_confirmation",
	})
	if err != nil {
		return
	}
	if validator.Fails() {
		ctx.Response().Json(http.StatusUnprocessableEntity, validator.Errors().All())
	}
	if err := helpers.PasswordReset(ctx); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "password reset successfully",
	})
}

func (r *PasswordController) Update(ctx http.Context) {
}
