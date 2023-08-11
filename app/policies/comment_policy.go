package policies

import (
	"context"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/auth/access"
)

type CommentPolicy struct {
}

func NewCommentPolicy() *CommentPolicy {
	return &CommentPolicy{}
}

func (r *CommentPolicy) Update(ctx context.Context, arguments map[string]any) access.Response {
	user := ctx.Value("user").(models.User)
	comment := arguments["comment"].(models.Comment)

	if user.ID == comment.UserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this comment.")
	}
}

func (r *CommentPolicy) Delete(ctx context.Context, arguments map[string]any) access.Response {
	user := ctx.Value("user").(models.User)
	comment := arguments["comment"].(models.Comment)

	if user.ID == comment.UserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("you do not own this comment.")
	}
}
