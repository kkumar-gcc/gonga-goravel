package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type CommentController struct {
	//Dependent services
}

func NewCommentController() *CommentController {
	return &CommentController{
		//Inject services
	}
}

func (r *CommentController) Index(ctx http.Context) {
	postID := ctx.Request().Route("id")
	var comments []models.Comment
	var total int64
	// Get page and per_page parameters from the request
	page := ctx.Request().QueryInt("page", 1)
	perPage := ctx.Request().QueryInt("per_page", 10)
	if err := facades.Orm().Query().Where("post_id = ? AND parent_id IS NULL", postID).Paginate(page, perPage, &comments, &total); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": comments,
		"meta": http.Json{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"last_page":    total / int64(perPage),
		},
	})
}

func (r *CommentController) Show(ctx http.Context) {
	id := ctx.Request().Route("id")
	var comment models.Comment
	if err := facades.Orm().Query().With("Children").Where("id = ?", id).FirstOrFail(&comment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": comment,
	})
}

func (r *CommentController) Store(ctx http.Context) {
	id := ctx.Request().RouteInt("id")
	var storeCommentRequest requests.StoreCommentRequest
	errors, err := ctx.Request().ValidateRequest(&storeCommentRequest)
	if err != nil || errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}
	if storeCommentRequest.ParentID != nil {
		var parentComment models.Comment
		if err := facades.Orm().Query().Where("id = ?", storeCommentRequest.ParentID).FirstOrFail(&parentComment); err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	var user models.User
	if err := facades.Auth().User(ctx, &user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	newComment := models.Comment{
		Body:     storeCommentRequest.Body,
		ParentID: storeCommentRequest.ParentID,
		PostID:   uint(id),
		UserID:   user.ID,
	}

	if err := facades.Orm().Query().Create(&newComment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(storeCommentRequest.Mentions) > 0 {
		for _, mention := range storeCommentRequest.Mentions {
			mention.OwnerID = newComment.ID
			mention.OwnerType = "comments"
			if err := facades.Orm().Query().Create(&mention); err != nil {
				helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	ctx.Response().Success().Json(http.Json{
		"data": newComment,
	})
}

func (r *CommentController) Update(ctx http.Context) {
	id := ctx.Request().Route("id")
	var updateCommentRequest requests.UpdateCommentRequest
	errors, err := ctx.Request().ValidateRequest(&updateCommentRequest)
	if err != nil || errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}
	var comment models.Comment
	if err := facades.Orm().Query().Where("id = ?", id).FirstOrFail(&comment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	comment.Body = updateCommentRequest.Body
	if err := facades.Orm().Query().Save(&comment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// TODO: Write logic for updating mentions in comment
	ctx.Response().Success().Json(http.Json{
		"data":    comment,
		"message": "comment updated successfully",
	})
}

func (r *CommentController) Delete(ctx http.Context) {
	id := ctx.Request().Route("id")
	var comment models.Comment
	if err := facades.Orm().Query().Where("id = ?", id).FirstOrFail(&comment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if _, err := facades.Orm().Query().Delete(&comment); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"message": "comment deleted successfully",
	})
}
