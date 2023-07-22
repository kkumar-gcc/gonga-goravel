package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/support/carbon"
)

type UpdateUserRequest struct {
	FirstName          string          `form:"first_name" json:"first_name"`
	LastName           string          `form:"last_name" json:"last_name"`
	AvatarURL          string          `form:"avatar_url" json:"avatar_url"`
	Bio                string          `form:"bio" json:"bio"`
	Gender             string          `form:"gender" json:"gender"`
	MobileNo           string          `form:"mobile_no" json:"mobile_no"`
	Country            string          `form:"country" json:"country"`
	City               string          `form:"city" json:"city"`
	MobileNoCode       string          `form:"mobile_no_code" json:"mobile_no_code"`
	Birthday           carbon.DateTime `form:"birthday" json:"birthday"`
	BackgroundImageURL string          `form:"background_image_url" json:"background_image_url"`
	WebsiteURL         string          `form:"website_url" json:"website_url"`
	Occupation         string          `form:"occupation" json:"occupation"`
	Education          string          `form:"education" json:"education"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"first_name":  "required,max_len:50",
		"last_name":   "required,max_len:50",
		"bio":         "max_len:500",
		"mobile_no":   "len:10",
		"country":     "alpha_dash,max_len:50",
		"city":        "alpha_dash,max_len:50",
		"birthday":    "date",
		"website_url": "full_url",
		"occupation":  "max_len:50",
		"education":   "max_len:50",
	}
}

func (r *UpdateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
