package models

import (
	"github.com/goravel/framework/database/orm"
)

type Mention struct {
	orm.Model
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	OwnerID   uint   `json:"owner_id"`
	OwnerType string `json:"owner_type"` // posts, comments, etc.
	Position  int    `json:"position"`   // position of the mention in the content
	orm.SoftDeletes
}
