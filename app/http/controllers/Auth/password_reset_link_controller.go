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
