package policies

import (
	"context"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/auth/access"
)

type UserPolicy struct {
}

func NewUserPolicy() *UserPolicy {
	return &UserPolicy{}
}

func (r *UserPolicy) Update(ctx context.Context, arguments map[string]any) access.Response {
	authUser := ctx.Value("user").(models.User)
	user := arguments["user"].(models.User)

	if authUser.ID == user.ID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this user account.")
	}
}

func (r *UserPolicy) Delete(ctx context.Context, arguments map[string]any) access.Response {
	authUser := ctx.Value("user").(models.User)
	user := arguments["user"].(models.User)

	if authUser.ID == user.ID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this user account.")
	}
}
