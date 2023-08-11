package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type PostController struct {
	//Dependent services
}

func NewPostController() *PostController {
	return &PostController{
		//Inject services
	}
}

func (r *PostController) Index(ctx http.Context) {
	var posts []models.Post
	var total int64
	// Get page and per_page parameters from the request
	page := ctx.Request().QueryInt("page", 1)
	perPage := ctx.Request().QueryInt("per_page", 10)
	if err := facades.Orm().Query().Omit("password").Paginate(page, perPage, &posts, &total); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": posts,
		"meta": http.Json{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"last_page":    total / int64(perPage),
		},
	})
}

func (r *PostController) Show(ctx http.Context) {
	id := ctx.Request().Route("id")
	var post models.Post
	if err := facades.Orm().Query().Where("id", id).FirstOrFail(&post); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": post,
	})
}

func (r *PostController) Store(ctx http.Context) {
	var storePostRequest requests.StorePostRequest
	errors, err := ctx.Request().ValidateRequest(&storePostRequest)
	if err != nil || errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}
	var user models.User
	if err := facades.Auth().User(ctx, &user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	newPost := models.Post{
		Title:              storePostRequest.Title,
		Body:               storePostRequest.Body,
		IsPromoted:         storePostRequest.IsPromoted,
		PromotionExpiresAt: storePostRequest.PromotionExpiresAt,
		IsFeatured:         storePostRequest.IsFeatured,
		FeaturedExpiresAt:  storePostRequest.FeaturedExpiresAt,
		UserID:             user.ID,
	}
	if err := facades.Orm().Query().Create(&newPost); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Associate the media files with the post
	for _, newMedia := range storePostRequest.Medias {
		_, err := facades.Orm().Query().Model(&models.Media{}).Where("id", newMedia.ID).Update(models.Media{
			OwnerID:   newPost.ID,
			OwnerType: "posts",
		})
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Iterate over the mention user IDs
	for _, mention := range storePostRequest.Mentions {
		err := facades.Orm().Query().Create(&models.Mention{
			UserID:    mention.UserID,
			OwnerID:   newPost.ID,
			OwnerType: "posts",
		})
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Create a slice to store the tags
	var tags []models.Tag

	for _, hashtag := range storePostRequest.Hashtags {
		tag := models.Tag{
			Title: hashtag.Title,
		}
		if err := facades.Orm().Query().FirstOrCreate(&tag); err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		tags = append(tags, tag)
	}

	if err := facades.Orm().Query().Model(&newPost).Association("Hashtags").Replace(tags); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"type":    "success",
		"message": "the post was created successfully",
	})
}

func (r *PostController) Update(ctx http.Context) {
}

func (r *PostController) UpdateTitle(ctx http.Context) {
}

func (r *PostController) UpdateBody(ctx http.Context) {
}

func (r *PostController) UpdateMedia(ctx http.Context) {
}

func (r *PostController) UpdateHashtag(ctx http.Context) {
}

func (r *PostController) UpdatePostSettings(ctx http.Context) {
}

func (r *PostController) Delete(ctx http.Context) {
	id := ctx.Request().Route("id")
	var post models.Post
	if err := facades.Orm().Query().Where("id", id).FirstOrFail(&post); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// Check if the authenticated user is the owner of the post
	response := facades.Gate().Inspect("delete-post", map[string]any{
		"post": post,
	})
	if !response.Allowed() {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, response.Message())
		return
	}

	if _, err := facades.Orm().Query().Delete(&post); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"type":    "success",
		"message": "the post was deleted successfully",
	})
}
