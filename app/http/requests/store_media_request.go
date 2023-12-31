package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"mime/multipart"
)

type StoreMediaRequest struct {
	Files     []multipart.FileHeader `form:"files" json:"files"`
	OwnerID   uint                   `form:"owner_id" json:"owner_id"`
	OwnerType string                 `form:"owner_type" json:"owner_type"`
}

func (r *StoreMediaRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreMediaRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"files": "required",
	}
}

func (r *StoreMediaRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreMediaRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreMediaRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
