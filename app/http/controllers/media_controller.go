package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type MediaController struct {
	//Dependent services
}

func NewMediaController() *MediaController {
	return &MediaController{
		//Inject services
	}
}

func (r *MediaController) Index(ctx http.Context) {
}
