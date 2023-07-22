package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/support/carbon"
	"goravel/app/models"
)

type StorePostRequest struct {
	Title              string            `form:"title" json:"title"`
	Body               string            `form:"body" json:"body"`
	Hashtags           []models.Tag      `form:"hashtags" json:"hashtags"`
	Mentions           []models.Mention  `form:"mentions" json:"mentions"`
	Medias             []models.Media    `form:"medias" json:"medias"`
	IsPromoted         bool              `form:"is_promoted" json:"is_promoted"`
	PromotionExpiresAt carbon.DateTime   `form:"promotion_expires_at" json:"promotion_expires_at"`
	IsFeatured         bool              `form:"is_featured" json:"is_featured"`
	FeaturedExpiresAt  carbon.DateTime   `form:"featured_expires_at" json:"featured_expires_at"`
	Visibility         models.Visibility `form:"visibility" json:"visibility"`
}

func (r *StorePostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StorePostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title":    "required,min_len:20",
		"body":     "required,min_len:40",
		"hashtags": "min_len:1,max_len:15",
		"mentions": "max_len:15",
		"medias":   "max_len:15",
	}
}

func (r *StorePostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StorePostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StorePostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
