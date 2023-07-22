package Auth

import (
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type EmailVerificationNotificationController struct {
	//Dependent services
}

func NewEmailVerificationNotificationController() *EmailVerificationNotificationController {
	return &EmailVerificationNotificationController{
		//Inject services
	}
}

func (r *EmailVerificationNotificationController) Index(ctx http.Context) {
}

func (r *EmailVerificationNotificationController) Store(ctx http.Context) {
	var user2 models.User
	err := facades.Auth().User(ctx, &user2) // Must point
	color.Redln(user2.Username, user2.Email, user2.Email)
	if err != nil {
		return
	}
}
