package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/support/carbon"
	"goravel/app/models"
)

type UpdatePostRequest struct {
	IsPromoted         bool              `form:"is_promoted" json:"is_promoted"`
	PromotionExpiresAt carbon.DateTime   `form:"promotion_expires_at" json:"promotion_expires_at"`
	IsFeatured         bool              `form:"is_featured" json:"is_featured"`
	FeaturedExpiresAt  carbon.DateTime   `form:"featured_expires_at" json:"featured_expires_at"`
	Visibility         models.Visibility `form:"visibility" json:"visibility"`
}

func (r *UpdatePostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatePostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"is_promoted":          "required",
		"promotion_expires_at": "required",
		"is_featured":          "required",
		"featured_expires_at":  "required",
		"visibility":           "required",
	}
}

func (r *UpdatePostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatePostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatePostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
