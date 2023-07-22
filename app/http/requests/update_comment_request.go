package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"goravel/app/models"
)

type UpdateCommentRequest struct {
	Body     string           `form:"body" json:"body"`
	Mentions []models.Mention `form:"mentions" json:"mentions"`
}

func (r *UpdateCommentRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateCommentRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"body":     "required,min_len:40",
		"mentions": "max_len:15",
	}
}

func (r *UpdateCommentRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateCommentRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateCommentRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
