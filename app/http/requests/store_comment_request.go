package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"goravel/app/models"
)

type StoreCommentRequest struct {
	Body     string           `form:"body" json:"body"`
	ParentID *uint            `form:"parent_id" json:"parent_id"`
	Mentions []models.Mention `form:"mentions" json:"mentions"`
}

func (r *StoreCommentRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreCommentRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"body":     "required",
		"mentions": "max_len:15",
	}
}

func (r *StoreCommentRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreCommentRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreCommentRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
