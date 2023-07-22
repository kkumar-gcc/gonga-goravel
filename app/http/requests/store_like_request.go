package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreLikeRequest struct {
	LikeableID   uint   `form:"likeable_id" json:"likeable_id"`
	LikeableType string `form:"likeable_type" json:"likeable_type"`
}

func (r *StoreLikeRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreLikeRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"likeable_id":   "required",
		"likeable_type": "required",
	}
}

func (r *StoreLikeRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreLikeRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreLikeRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
