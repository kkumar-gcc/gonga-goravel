package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type NotificationController struct {
	//Dependent services
}

func NewNotificationController() *NotificationController {
	return &NotificationController{
		//Inject services
	}
}

func (r *NotificationController) Index(ctx http.Context) {
}

func (r *NotificationController) ReadAll(ctx http.Context) {
}

func (r *NotificationController) Update(ctx http.Context) {
}
