package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type LikeController struct {
	//Dependent services
}

func NewLikeController() *LikeController {
	return &LikeController{
		//Inject services
	}
}

func (r *LikeController) Store(ctx http.Context) {
	var storeLikeRequest requests.StoreLikeRequest
	errors, err := ctx.Request().ValidateRequest(&storeLikeRequest)
	if err != nil || errors != nil {
		ctx.Response().Json(http.StatusUnprocessableEntity, errors.All())
		return
	}
	var user models.User
	if err := facades.Auth().User(ctx, &user); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	var count int64
	if err := facades.Orm().Query().Table(storeLikeRequest.LikeableType).Where("id = ?", storeLikeRequest.LikeableID).Count(&count); err != nil || count == 0 {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	var like models.Like
	like.UserID = user.ID
	like.LikeableID = storeLikeRequest.LikeableID
	like.LikeableType = storeLikeRequest.LikeableType
	// check if the user has already liked this likeable
	if err := facades.Orm().Query().FirstOrFail(&like); err == nil {
		// User has already liked this likeable, perform unlike operation
		if _, err := facades.Orm().Query().Delete(&like); err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{
				"error": err.Error(),
			})
			return
		}
		ctx.Response().Success().Json(http.Json{
			"message": "like deleted successfully",
		})
		return
	}
	if err := facades.Orm().Query().Create(&like); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	ctx.Response().Success().Json(http.Json{
		"message": "like created successfully",
	})
}

func (r *LikeController) Delete(ctx http.Context) {
	id := ctx.Request().Route("id")
	var like models.Like
	if err := facades.Orm().Query().Where("id = ?", id).FirstOrFail(&like); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if err := facades.Auth().User(ctx, &user); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	if like.UserID != user.ID {
		ctx.Response().Json(http.StatusForbidden, http.Json{
			"error": "you are not authorized to delete this like",
		})
		return
	}
	if _, err := facades.Orm().Query().Delete(&like); err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	ctx.Response().Success().Json(http.Json{
		"message": "like deleted successfully",
	})
}
