package models

import (
	"github.com/goravel/framework/database/orm"
)

type PasswordReset struct {
	orm.Model
	Email  string `gorm:"unique;not null"`
	Token  string `gorm:"not null"`
	Expiry int64  `gorm:"not null"`
	orm.SoftDeletes
}
