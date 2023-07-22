package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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
		"password": "required|min_len:8",
	})
	if err != nil {
		return
	}
	if validator.Fails() {
		ctx.Response().Json(http.StatusUnprocessableEntity, validator.Errors().All())
	}
	//token := ctx.Request().Query("token")
	email := ctx.Request().Query("email")
	password := ctx.Request().Query("password")

	_, err = facades.Orm().Query().Where("email", email).Update(map[string]interface{}{
		"password": facades.Hash().Make(password),
	})
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

}

func (r *PasswordController) Update(ctx http.Context) {
}
