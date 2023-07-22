package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdatePostRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *UpdatePostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatePostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
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
