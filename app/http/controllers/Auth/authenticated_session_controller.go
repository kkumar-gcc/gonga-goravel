package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	Auth2 "goravel/app/http/requests/Auth"
	"goravel/app/models"
)

type AuthenticatedSessionController struct {
	//Dependent services
}

func NewAuthenticatedSessionController() *AuthenticatedSessionController {
	return &AuthenticatedSessionController{
		//Inject services
	}
}

func (r *AuthenticatedSessionController) Index(ctx http.Context) {
}

func (r *AuthenticatedSessionController) Store(ctx http.Context) {
	var loginRequest Auth2.LoginRequest
	errors, err := ctx.Request().ValidateRequest(&loginRequest)
	if err != nil || errors != nil {
		ctx.Response().Json(http.StatusUnprocessableEntity, errors.All())
		return
	}

	var user models.User
	if err := facades.Orm().Query().Where("username", loginRequest.Username).First(&user); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	if !facades.Hash().Check(loginRequest.Password, user.Password) {
		ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := facades.Auth().Login(ctx, &user)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	ctx.Response().Header("Authorization", token).Status(http.StatusNoContent)
}

func (r *AuthenticatedSessionController) Destroy(ctx http.Context) {
	if err := facades.Auth().Logout(ctx); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	ctx.Response().Status(http.StatusNoContent)
}
