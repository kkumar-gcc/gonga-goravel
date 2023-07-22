package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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
		ctx.Response().Json(http.StatusUnprocessableEntity, errors.All())
		return
	}

	hashedPassword, _ := facades.Hash().Make(registerRequest.Password)
	var user models.User
	user.Username = registerRequest.Username
	user.Password = hashedPassword
	user.Email = registerRequest.Email
	if err := facades.Orm().Query().Create(&user); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	token, err := facades.Auth().Login(ctx, &user)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	ctx.Response().Header("Authorization", token).Status(http.StatusNoContent)
}
