package Auth

import (
	"github.com/goravel/framework/contracts/http"
)

type NewPasswordController struct {
	//Dependent services
}

func NewNewPasswordController() *NewPasswordController {
	return &NewPasswordController{
		//Inject services
	}
}

func (r *NewPasswordController) Index(ctx http.Context) {
}
