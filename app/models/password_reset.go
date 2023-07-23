package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type PasswordReset struct {
	orm.Model
	Email     string          `gorm:"unique;not null"`
	Token     string          `gorm:"not null"`
	ExpiresAt carbon.DateTime `gorm:"not null"`
}
