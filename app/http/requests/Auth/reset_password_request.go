package Auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ResetPasswordRequest struct {
	Email string `form:"email" json:"email"`
}

func (r *ResetPasswordRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ResetPasswordRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "required,email",
	}
}

func (r *ResetPasswordRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ResetPasswordRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ResetPasswordRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
