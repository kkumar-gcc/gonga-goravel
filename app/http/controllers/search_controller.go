package controllers

import (
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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
	color.Redln("SearchController.Index", facades.Config().GetString("cloudinary.cloud_name"))
}
