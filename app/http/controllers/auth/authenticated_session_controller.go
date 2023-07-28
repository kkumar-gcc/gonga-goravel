package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
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
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	var user models.User
	if err := facades.Orm().Query().Where("username", loginRequest.Username).First(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !facades.Hash().Check(loginRequest.Password, user.Password) {
		helpers.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := facades.Auth().Login(ctx, &user)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Header("Authorization", token).Status(http.StatusNoContent)
}

func (r *AuthenticatedSessionController) Destroy(ctx http.Context) {
	if err := facades.Auth().Logout(ctx); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Status(http.StatusNoContent)
}
