package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Index(ctx http.Context) {
}

func (r *UserController) Show(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Create(ctx http.Context) {
}

func (r *UserController) Update(ctx http.Context) {
}

func (r *UserController) Delete(ctx http.Context) {
}
