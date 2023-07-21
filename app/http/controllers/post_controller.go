package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type PostController struct {
	//Dependent services
}

func NewPostController() *PostController {
	return &PostController{
		//Inject services
	}
}

func (r *PostController) Index(ctx http.Context) {
}
