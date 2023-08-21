package policies

import (
	"context"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/auth/access"
)

type PostPolicy struct {
}

func NewPostPolicy() *PostPolicy {
	return &PostPolicy{}
}

func (r *PostPolicy) Create(ctx context.Context, arguments map[string]any) access.Response {
	return nil
}

func (r *PostPolicy) Update(ctx context.Context, arguments map[string]any) access.Response {
	user := ctx.Value("user").(models.User)
	post := arguments["post"].(models.Post)

	if user.ID == post.UserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this post.")
	}
}

func (r *PostPolicy) Delete(ctx context.Context, arguments map[string]any) access.Response {
	user := ctx.Value("user").(models.User)
	post := arguments["post"].(models.Post)

	if user.ID == post.UserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this post.")
	}
}
