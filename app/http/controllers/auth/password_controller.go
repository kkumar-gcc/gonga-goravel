package auth

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
		"email":    "required|email",
		"password": "required|min_len:8|eq_field:password_confirmation",
	})
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if validator.Fails() {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, validator.Errors().All())
		return
	}
	if err := helpers.PasswordReset(ctx); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "password reset successfully",
	})
}

func (r *PasswordController) Update(ctx http.Context) {
}
