package config

import (
	cloudinaryFacades "github.com/goravel/cloudinary/facades"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("filesystems", map[string]any{
		// Default Filesystem Disk
		//
		// Here you may specify the default filesystem disk that should be used
		// by the framework. The "local" disk, as well as a variety of cloud
		// based disks are available to your application. Just store away!
		"default": config.Env("FILESYSTEM_DISK", "local"),

		// Filesystem Disks
		//
		// Here you may configure as many filesystem "disks" as you wish, and you
		// may even configure multiple disks of the same driver. Defaults have
		// been set up for each driver as an example of the required values.
		//
		// Supported Drivers: "local", "s3", "oss", "cos", "minio", "custom"
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   "storage/app",
				"url":    config.Env("APP_URL", "").(string) + "/storage",
			},
			"cloudinary": map[string]any{
				"driver": "custom",
				"cloud":  config.Env("CLOUDINARY_CLOUD_NAME", ""),
				"key":    config.Env("CLOUDINARY_API_KEY", ""),
				"secret": config.Env("CLOUDINARY_API_SECRET", ""),
				"via": func() (filesystem.Driver, error) {
					return cloudinaryFacades.Cloudinary("cloudinary"), nil
				},
			},
		},
	})
}
