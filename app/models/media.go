package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"

	"goravel/database/factories"
)

type Media struct {
	orm.Model
	URL       string `json:"url"`
	Type      string `json:"type"`
	OwnerID   uint   `json:"owner_id"`
	OwnerType string `json:"owner_type"` // posts, comments, users, etc.
	orm.SoftDeletes
}

func (m *Media) Factory() factory.Factory {
	return &factories.MediaFactory{}
}
