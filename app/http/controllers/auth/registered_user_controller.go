package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	Auth2 "goravel/app/http/requests/Auth"
	"goravel/app/models"
)

type RegisteredUserController struct {
	//Dependent services
}

func NewRegisteredUserController() *RegisteredUserController {
	return &RegisteredUserController{
		//Inject services
	}
}

func (r *RegisteredUserController) Index(ctx http.Context) {
}

func (r *RegisteredUserController) Store(ctx http.Context) {
	var registerRequest Auth2.RegisterRequest
	errors, err := ctx.Request().ValidateRequest(&registerRequest)
	if err != nil || errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	hashedPassword, err := facades.Hash().Make(registerRequest.Password)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var user models.User
	user.Username = registerRequest.Username
	user.Password = hashedPassword
	user.Email = registerRequest.Email
	if err := facades.Orm().Query().Create(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := facades.Auth().Login(ctx, &user)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Header("Authorization", token).Status(http.StatusNoContent)
}
