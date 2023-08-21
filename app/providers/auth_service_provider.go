package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"goravel/app/policies"
)

type AuthServiceProvider struct {
}

func (receiver *AuthServiceProvider) Register(app foundation.Application) {

}

func (receiver *AuthServiceProvider) Boot(app foundation.Application) {
	facades.Gate().Define("update-post", policies.NewPostPolicy().Update)
	facades.Gate().Define("delete-post", policies.NewPostPolicy().Delete)
	facades.Gate().Define("update-comment", policies.NewCommentPolicy().Update)
	facades.Gate().Define("delete-comment", policies.NewCommentPolicy().Delete)
	facades.Gate().Define("update-user", policies.NewUserPolicy().Update)
	facades.Gate().Define("delete-user", policies.NewUserPolicy().Delete)
}
