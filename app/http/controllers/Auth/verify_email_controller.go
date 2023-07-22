package Auth

import (
	"github.com/goravel/framework/contracts/http"
)

type VerifyEmailController struct {
	//Dependent services
}

func NewVerifyEmailController() *VerifyEmailController {
	return &VerifyEmailController{
		//Inject services
	}
}

func (r *VerifyEmailController) Index(ctx http.Context) {
}
