package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type FollowController struct {
	//Dependent services
}

func NewFollowController() *FollowController {
	return &FollowController{
		//Inject services
	}
}

func (r *FollowController) Index(ctx http.Context) {
}
