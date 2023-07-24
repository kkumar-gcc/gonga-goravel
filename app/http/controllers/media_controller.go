package controllers

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type MediaController struct {
	//Dependent services
}

func NewMediaController() *MediaController {
	return &MediaController{
		//Inject services
	}
}

var results []interface{}

func (r *MediaController) Store(ctx http.Context) {
	var storeMediaRequest requests.StoreMediaRequest
	errors, err := ctx.Request().ValidateRequest(&storeMediaRequest)
	if err != nil || errors != nil {
		color.Redln("MediaController.Store", err, errors)
		ctx.Response().Json(http.StatusUnprocessableEntity, errors.All())
		return
	}
	files := storeMediaRequest.Files
	ownerID := storeMediaRequest.OwnerID
	ownerType := storeMediaRequest.OwnerType
	cloudName := facades.Config().GetString("cloudinary.cloud_name")
	apiKey := facades.Config().GetString("cloudinary.api_key")
	apiSecret := facades.Config().GetString("cloudinary.api_secret")
	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{
				"error": err.Error(),
			})
			return
		}

		publicID := "test"
		// Upload image
		upload, err := cld.Upload.Upload(ctx.Context(), file, uploader.UploadParams{
			PublicID:       publicID,
			UniqueFilename: api.Bool(false),
			Overwrite:      api.Bool(true),
		})
		if err != nil {
			return
		}
		// Close the file explicitly here, after the upload is done or in case of an error
		if err := file.Close(); err != nil {
			return
		}
		// Create media
		media := models.Media{
			OwnerID:   ownerID,
			OwnerType: ownerType,
			URL:       upload.URL,
			Type:      upload.Type,
		}
		if err := facades.Orm().Query().Create(&media); err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{
				"error": err.Error(),
			})
			return
		}

		result := map[string]interface{}{
			"id":        media.ID,
			"url":       upload.URL,
			"type":      upload.Type,
			"file_name": upload.OriginalFilename,
			"size":      upload.Bytes,
		}
		results = append(results, result)
	}
	ctx.Response().Success().Json(http.Json{
		"message": "files uploaded successfully",
		"data":    results,
	})
}
