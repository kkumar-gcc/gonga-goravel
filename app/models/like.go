package models

import (
	"github.com/goravel/framework/database/orm"
)

type Like struct {
	orm.Model
	UserID       uint   `json:"user_id"`
	User         *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	LikeableID   uint   `json:"likable_id"`
	LikeableType string `json:"likable_type"` // posts, comments, users, etc.
	orm.SoftDeletes
}
