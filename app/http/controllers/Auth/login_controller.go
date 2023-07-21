package Auth

import (
	"github.com/goravel/framework/contracts/http"
)

type LoginController struct {
	//Dependent services
}

func NewLoginController() *LoginController {
	return &LoginController{
		//Inject services
	}
}

func (r *LoginController) Index(ctx http.Context) {
}

func (r *LoginController) Create(ctx http.Context) {
}

func (r *LoginController) Delete(ctx http.Context) {
}
