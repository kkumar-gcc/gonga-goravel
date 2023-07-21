package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type CommentController struct {
	//Dependent services
}

func NewCommentController() *CommentController {
	return &CommentController{
		//Inject services
	}
}

func (r *CommentController) Index(ctx http.Context) {
}
