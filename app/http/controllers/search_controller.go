package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type SearchController struct {
	//Dependent services
}

func NewSearchController() *SearchController {
	return &SearchController{
		//Inject services
	}
}

func (r *SearchController) Index(ctx http.Context) {
}
