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

func (r *PostController) Show(ctx http.Context) {
}

func (r *PostController) Store(ctx http.Context) {
}

func (r *PostController) Update(ctx http.Context) {
}

func (r *PostController) UpdateTitle(ctx http.Context) {
}

func (r *PostController) UpdateBody(ctx http.Context) {
}

func (r *PostController) UpdateMedia(ctx http.Context) {
}

func (r *PostController) UpdateHashtag(ctx http.Context) {
}

func (r *PostController) UpdatePostSettings(ctx http.Context) {
}

func (r *PostController) Delete(ctx http.Context) {
}
