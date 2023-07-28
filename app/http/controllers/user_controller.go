package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Index(ctx http.Context) {
	var users []models.User
	var total int64
	// Get page and per_page parameters from the request
	page := ctx.Request().QueryInt("page", 1)
	perPage := ctx.Request().QueryInt("per_page", 10)
	if err := facades.Orm().Query().Omit("password").Paginate(page, perPage, &users, &total); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": users,
		"meta": http.Json{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"last_page":    total / int64(perPage),
		},
	})
}

func (r *UserController) Show(ctx http.Context) {
	username := ctx.Request().Route("username")
	var user models.User
	if err := facades.Orm().Query().Where("username", username).Omit("password").FirstOrFail(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": user,
	})
}

func (r *UserController) Update(ctx http.Context) {
	username := ctx.Request().Route("username")
	var updateUserRequest requests.UpdateUserRequest
	errors, err := ctx.Request().ValidateRequest(&updateUserRequest)
	if err != nil || errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}
	var user models.User
	if err := facades.Orm().Query().Where("username", username).FirstOrFail(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if updateUserRequest.FirstName != "" {
		user.FirstName = updateUserRequest.FirstName
	}
	if updateUserRequest.LastName != "" {
		user.LastName = updateUserRequest.LastName
	}
	if updateUserRequest.AvatarURL != "" {
		user.AvatarURL = updateUserRequest.AvatarURL
	}
	if updateUserRequest.Bio != "" {
		user.Bio = updateUserRequest.Bio
	}
	if updateUserRequest.Gender != "" {
		user.Gender = updateUserRequest.Gender
	}
	if updateUserRequest.MobileNo != "" {
		user.MobileNo = updateUserRequest.MobileNo
	}
	if updateUserRequest.MobileNoCode != "" {
		user.MobileNoCode = updateUserRequest.MobileNoCode
	}
	if updateUserRequest.Country != "" {
		user.Country = updateUserRequest.Country
	}
	if updateUserRequest.City != "" {
		user.City = updateUserRequest.City
	}
	if updateUserRequest.Birthday.IsZero() {
		user.Birthday = updateUserRequest.Birthday
	}
	if updateUserRequest.BackgroundImageURL != "" {
		user.BackgroundImageURL = updateUserRequest.BackgroundImageURL
	}
	if updateUserRequest.WebsiteURL != "" {
		user.WebsiteURL = updateUserRequest.WebsiteURL
	}
	if updateUserRequest.Occupation != "" {
		user.Occupation = updateUserRequest.Occupation
	}
	if updateUserRequest.Education != "" {
		user.Education = updateUserRequest.Education
	}
	if err := facades.Orm().Query().Model(&models.User{}).Save(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": user,
	})
}

func (r *UserController) Delete(ctx http.Context) {
	id := ctx.Request().Route("id")
	var user models.User
	if err := facades.Orm().Query().Where("id", id).FirstOrFail(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if _, err := facades.Orm().Query().Model(&models.User{}).Delete(&user); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Response().Success().Json(http.Json{
		"data": user,
	})
}
