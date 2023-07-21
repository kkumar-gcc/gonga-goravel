package models

import (
	"github.com/goravel/framework/database/orm"
)

type Follow struct {
	orm.Model
	FollowerID  uint `json:"follower_id" gorm:"not null"`
	FollowingID uint `json:"following_id" gorm:"not null"`
	orm.SoftDeletes
}
