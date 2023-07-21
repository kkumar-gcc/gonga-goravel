package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type LikeController struct {
	//Dependent services
}

func NewLikeController() *LikeController {
	return &LikeController{
		//Inject services
	}
}

func (r *LikeController) Index(ctx http.Context) {
}
