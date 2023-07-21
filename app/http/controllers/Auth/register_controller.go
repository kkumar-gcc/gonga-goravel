package Auth

import (
	"github.com/goravel/framework/contracts/http"
)

type RegisterController struct {
	//Dependent services
}

func NewRegisterController() *RegisterController {
	return &RegisterController{
		//Inject services
	}
}

func (r *RegisterController) Index(ctx http.Context) {
}
