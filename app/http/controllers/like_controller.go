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

func (r *LikeController) Create(ctx http.Context) {
}

func (r *LikeController) Delete(ctx http.Context) {
}
