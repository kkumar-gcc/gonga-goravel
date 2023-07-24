package config

import "github.com/goravel/framework/facades"

func init() {
	config := facades.Config()
	config.Add("cloudinary", map[string]any{
		// Cloudinary Configuration
		//
		// Here you may configure your settings for cloudinary.
		//
		// To learn more: https://cloudinary.com/documentation/go_integration
		"cloud_name": config.Env("CLOUDINARY_CLOUD_NAME", ""),
		"api_key":    config.Env("CLOUDINARY_API_KEY", ""),
		"api_secret": config.Env("CLOUDINARY_API_SECRET", ""),
	})
}
